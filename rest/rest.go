package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jhin93/Golang_nomadCoin/blockchain"
	"github.com/jhin93/Golang_nomadCoin/utils"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"` // struct field tag 작성방법. struct field를 소문자로 json으로 보여준다.
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty는 해당 field의 value가 비어있을 경우, 해당 field 자체를 생략해준다.
}

type addBlockBody struct {
	Message string // request로 들어오는 데이터의 key name("message")과 동일해야 한다.
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See All Blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}
	rw.Header().Add("Content-Type", "application/json") // text가 아닌 json으로 브라우저에 보내도록 Header에서 Content-type을 변경
	json.NewEncoder(rw).Encode(data)                    // .NewEncoder는 w에 작성할 encoder를 반환.
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks()) // 뭔가를 http client에 write할때 Encoder를 만들어서 encode 함.
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody)) // Decoder(.NewDecoder)를 만들어서 request(r)의 body(Body)로부터 읽어온다(.Decode). 그리고 그 결과를 실제(&) addBlockBody에 담는다(&addblockBody)
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated) // 에러처리를 위해 header에 http status 작성. StatusCreated는 201을 의미
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // mux.Vars()가 r에서 변수를 추출해줌
	// fmt.Println(vars) 결과 : map[id:1]
	id, err := strconv.Atoi(vars["height"]) // strconv 패키지의 Atoi 메소드는 string을 interger로 변환
	utils.HandleErr(err)
	block, err := blockchain.GetBlockchain().GetBlock(id) // id는 string이고, GetBlock()의 결과는 int이므로 타입 통일이 필요
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)}) // Sprint는 그냥 string으로 바꿔줌.
	} else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler { // middleware. middleware는 마지막 목적지 전에 호출됨.
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) { // HandlerFunc는 타입, http.Handler는 interface. 근데 왜 리턴타입이 매칭이 안되는데 에러가 안날까?
		// http.HandlerFunc는 'adapter'라고 불리는데 이는 인자로 들어간 함수가 조건에 부합하면 필요한 함수를 구현해준다. 그래서 위 'type url' 사례처럼 직접 url 혹은 type을 만들지 않아도 adapter가 대신 만들어줌. 이를 adapter 패턴이라 한다.
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)      // %d는 integer
	router := mux.NewRouter()             // gorilla mux 사용(mux.NewRouter())
	router.Use(jsonContentTypeMiddleware) // middleware가 먼저 호출 되고 next handler가 호출됨(next.ServeHTTP(rw, r)). 그건 아래 셋(router.HandleFunc()) 혹은 다른 middleware 중 하나가 될 것.
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET") // 변수({id})에 패턴 추가. mux가 확인. .Methods("GET") <- mux가 다른 method로부터 보호해줌. mux를 사용하지 않으면 위 switch문에서 .StatusMethodNotAllowed 같은 처리를 해주어야 함.
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router)) // 서버에게 기본 ServeMux(nil)이 아닌 커스텀 ServeMux(handler)사용한다고 말해줌.
}
