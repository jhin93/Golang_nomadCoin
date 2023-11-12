package main

import (
	"fmt"

	"github.com/jhin93/Golang_nomadCoin/person"
)

func main() {
	nico := person.Person{}
	nico.SetDetails("nico", 12)
	fmt.Println("Main's nico : ", nico)
}
