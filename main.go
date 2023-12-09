package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jhin93/Golang_nomadCoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string              // html 파일에 공유되어야 하기에 대문자
	Blocks    []*blockchain.Block // 블록을 렌더링하기 위해 선언. Declare as uppercase for sharing
}

// 첫번째 인자는 ResponseWriter. 이는 유저에게 보내고 싶은 데이터를 적는 것.
// 두번쨰 인자는 request pointer. request를 복사하려는 게 아니기에 포인터로 사용
func home(rw http.ResponseWriter, r *http.Request) {
	// 렌더링을 위해 html/template 메소드 사용. It return pointer and error(*template.Template, error).
	// template 메소드의 에러를 대신 처리해주는 메소드 .Must. 에러가 없다면 template pointer 반환(t *template.Template) 메소드 기능은 cmd+클릭으로 확인 가능.
	tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()} // 모든 블록체인의 블록 렌더링
	tmpl.Execute(rw, data)                                           // argument : 1. writer(wr io.Writer) 2. data(any)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	// .Fatal은 os.Exit(1) 다음에 따라나오는 error를 Print()
	// os.Exit(1)은 프로그램이 error code 1으로 종료되는 것
	log.Fatal(http.ListenAndServe(port, nil)) // ListenAndServe는 error를 반환하는 함수. 만약 에러가 있다면 log.Fatal발동. 아니라면 http.ListenAndServe 유지
}
