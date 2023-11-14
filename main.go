package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash)) // 1. byte 값 가져오기(data + prevHash) + sha256으로 난수화
	hexHash := fmt.Sprintf("%x", hash)                                       // 2. 16진수 string으로 변환(.Sprintf("%x"))
	genesisBlock.hash = hexHash                                              // 3. hash에 담음
	fmt.Println(genesisBlock)
}
