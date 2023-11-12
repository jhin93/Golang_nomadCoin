package main

import "fmt"

func main() {
	a := 2
	b := a
	a = 12
	fmt.Println(&b, &a)
}
