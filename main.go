package main

import (
	"github.com/jhin93/Golang_nomadCoin/explorer"
	"github.com/jhin93/Golang_nomadCoin/rest"
)

func main() {
	explorer.Start(3000)
	rest.Start(4000)
}
