package main

import "fmt"

func main() {
	foods := []string{"a", "b", "c"}
	fmt.Printf("%v\n", foods)
	foods = append(foods, "tomato")
	fmt.Printf("%v\n", foods)
}
