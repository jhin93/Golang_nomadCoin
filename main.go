package main

import (
	"fmt"

	"github.com/jhin93/Golang_nomadCoin/person"
)

func main() {
	nico := person.Person{}           // nico 변수에 'Person' type 할당
	nico.SetDetails("nico", 12)       // 결과 : SeeDetails nico : {nico 12}
	fmt.Println("Main's nico'", nico) // 결과 : Main's nico' { 0}
	// 결과 원인 : 변수 nico는 SetDetails 함수의 p 변수를 복사한 값이라 처음 복사됐을 때 빈 값 그대로({ 0}) 출력됨.
}
