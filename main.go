package main

import "fmt"

func main() {
	name := "Nicolas!! Is my name"
	for _, letter := range name {
		fmt.Printf("%x\n", letter)
	}
}
