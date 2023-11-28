package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

// 첫번째 인자는 ResponseWriter. 이는 유저에게 보내고 싶은 데이터를 적는 것.
// 두번쨰 인자는 request pointer. request를 복사하려는 게 아니기에 포인터로 사용
func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from home!") // Fprint는 콘솔이 아닌 Writer에 출력하는 것. data를 format해서 Writer에 보내는 것
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	// .Fatal은 os.Exit(1) 다음에 따라나오는 error를 Print()
	// os.Exit(1)은 프로그램이 error code 1으로 종료되는 것
	log.Fatal(http.ListenAndServe(port, nil)) // ListenAndServe는 error를 반환하는 함수. 만약 에러가 있다면 log.Fatal발동. 아니라면 http.ListenAndServe 유지
}
