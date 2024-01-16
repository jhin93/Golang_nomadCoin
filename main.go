package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jhin93/Golang_nomadCoin/blockchain"
)

const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL         URL    `json:"url"` // struct field tag 작성방법. struct field를 소문자로 json으로 보여준다.
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty는 해당 field의 value가 비어있을 경우, 해당 field 자체를 생략해준다.
}

type AddBlockBody struct {
	Message string // request로 들어오는 데이터의 key name("message")과 동일해야 한다.
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
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
		var addBlockBody AddBlockBody
		json.NewDecoder(r.Body).Decode(&addBlockBody) // Decoder(.NewDecoder)를 만들어서 request(r)의 body(Body)로부터 읽어온다(.Decode). 그리고 그 결과를 실제(&) addBlockBody에 담는다(&addblockBody)
		fmt.Println(addBlockBody)
	}
}

func main() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
