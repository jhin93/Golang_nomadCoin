package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jhin93/Golang_nomadCoin/utils"
)

const port string = ":4000"

type URLDescription struct {
	URL         string
	Method      string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	// Marshal()은 Go의 데이터를 json으로 변환. json과 에러를 반환. JSON을 Go의 데이터로 바꿀 땐 Unmarshal
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Printf("%s", b) // byte slice -> string
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
