package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer:	Start the HTML Explorer\n")
	fmt.Printf("rest:		Start the REST API (recommended)\n\n")
	// main의 switch문이 실행될 때, 두번째 argument(os.Args[1])가 없을 때를 대비해야 함. default가 출력되면 프로그램 종료해야 한다.
	os.Exit(0) // 그래서 아무 에러없이 프로그램 종료하는 명령어 사용(os.Exit(0))
}

func main() {
	if len(os.Args) < 2 { // os.Args는 'go run main.go' 뒤에 붙은 argument를 의미한다. Ex) go run main.go rest, go run main.go explorer
		// os.Args[0] 는 프로그램 이름이다. ex) [/var/folders/__/0g968kx52q16jwbpr6p825rc0000gn/T/go-build121872943/b001/exe/main
		usage()
	}
}
