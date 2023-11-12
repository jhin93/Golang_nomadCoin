package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) sayHello() { // type 'person'에만 유효한 함수(p person)
	fmt.Printf("Hello! my name is %s and my korean age is %d", p.name, p.age+1)
}

func main() {
	nico := person{name: "nico", age: 12}
	nico.sayHello()
}
