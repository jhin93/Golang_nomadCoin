package main

import "fmt"

func main() {
	a := 2
	b := &a
	fmt.Println(b, *b)
}

// 0xc000012088 2
