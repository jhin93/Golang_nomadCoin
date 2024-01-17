package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{id}"),
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

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort) // %d는 integer
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
