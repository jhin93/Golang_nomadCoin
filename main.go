package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) sayHello() { // 'person' struct의 메소드
	fmt.Printf("Hello! my name is %s and i'm %d", p.name, p.age) // .Println으로 출력하면 '...name is %s and i'm %d' 라고 출력됨.
	// %s, %d 는 각각 string, decimal(10진수)를 의미함. 순서대로 %s에는 p.name, %d에는 p.age가 할당됨
}

func main() {
	nico := person{name: "nico", age: 12}
	nico.sayHello()
}
