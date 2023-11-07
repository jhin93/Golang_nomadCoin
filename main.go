package main

import "fmt"

func main() {
	x := 412412
	xAsBinary := fmt.Sprintf("%b\n", x) // binary로 format 된 x의 string
	fmt.Println(x, xAsBinary)
}
