package main

import (
	"flag"
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
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest") // default value가 "rest"
	flag.Parse()

	fmt.Println(*port, *mode) // flag.Int 및 flag.String 함수는 각각 int와 string 타입의 포인터를 반환. 그래서 *port, *mode로 반환.
	// go run main.go
	// 결과 : 4000 rest
	// go run main.go -mode html -port 5000
	// 결과 : 5000 html
	// go run main.go -mode 4343 -port sdsd
	// 결과 :
	// invalid value "sdsd" for flag -port: parse error
	// Usage of /var/folders/__/0g968kx52q16jwbpr6p825rc0000gn/T/go-build919758010/b001/exe/main:
	//   -mode string
	// 		Choose between 'html' and 'rest (default "rest")
	//   -port int
	// 		Set port of the server (default 4000)
	// exit status 2
}
