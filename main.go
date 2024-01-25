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
}

func main() {
	if len(os.Args) < 2 { // os.Args는 'go run main.go' 뒤에 붙은 argument를 의미한다. Ex) go run main.go rest, go run main.go explorer
		usage()
	}
}
