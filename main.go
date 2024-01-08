package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL         string `json:"url"` // struct field tag 작성방법. struct field를 소문자로 json으로 보여준다.
	Method      string `json:"method"`
	Description string `json:"description"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	rw.Header().Add("Content-Type", "application/json") // text가 아닌 json으로 브라우저에 보내도록 Header에서 Content-type을 변경
	json.NewEncoder(rw).Encode(data)                    // .NewEncoder는 w에 작성할 encoder를 반환.
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
