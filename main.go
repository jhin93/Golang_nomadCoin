package main

import (
	"fmt"

	"github.com/jhin93/Golang_nomadCoin/person"
)

func main() {
	nico := person.Person{}             // nico 변수에 'Person' type 할당
	nico.SetDetails("nico", 12)         // 결과 : SeeDetails nico : {nico 12}
	fmt.Println("Main's nico : ", nico) // 결과 : Main's nico' { 0}
}
