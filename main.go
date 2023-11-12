package main

import "fmt"

func main() {
	a := 2
	b := a // a의 value를 복사한 것
	a = 12
	fmt.Println(b)
}
