package main

import "fmt"

func main() {
	a := 2
	b := &a // b의 타입은 int pointer. var b *int
	a = 50
	a = 12
	fmt.Println(*b)
}
