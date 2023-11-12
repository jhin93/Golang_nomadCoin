package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 50
	a = 12
	fmt.Println(*b)
}

// 12
