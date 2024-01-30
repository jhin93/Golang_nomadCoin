package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jhin93/Golang_nomadCoin/explorer"
	"github.com/jhin93/Golang_nomadCoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000:	Set the PORT of the server\n")
	fmt.Printf("-mode=rest:		Choose between 'html' and 'rest'\n\n")
	// main의 switch문이 실행될 때, 두번째 argument(os.Args[1])가 없을 때를 대비해야 함. default가 출력되면 프로그램 종료해야 한다.
	os.Exit(0) // 그래서 아무 에러없이 프로그램 종료하는 명령어 사용(os.Exit(0))
}

func main() {
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'") // default value가 "rest"
	flag.Parse()

	switch *mode {
	case "rest": // go run main.go -mode=rest -port=4000 [결과 : Listening on http://localhost:4000]
		rest.Start(*port)
	case "html": // go run main.go -mode=html [결과 : Listening on http://localhost:4000]
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode) // flag.Int 및 flag.String 함수는 각각 int와 string 타입의 포인터를 반환. 그래서 *port, *mode로 반환.

	// go run main.go -mode=xxx -port=4000
	// 결과 :
	// Welcome to 노마드 코인

	// Please use the following flags:

	// -port=4000:     Set the PORT of the server
	// -mode=rest:             Choose between 'html' and 'rest'
}
