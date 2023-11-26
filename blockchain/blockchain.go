package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*block // blocks는 pointer의 slice
}

var b *blockchain
var once sync.Once // sync 패키지의 type 'Once'

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash) // 16진수
}

func getLastHash() string { // 마지막 블록 해쉬 반환
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *block { // block 타입의 pointer 반환
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock // 반환값이 메모리 연산자인 이유는 createBlock 함수가 사용될 blockchain구조체의 내부 요소 blocks가 포인터 변수 슬라이스이기 때문
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain { // 변수 b와 동일한 타입인 blockchain의 pointer 반환
	// 1. b 변수 초기화 여부 확인
	if b == nil {
		once.Do(func() { // 내부의 로직을 오직 한번만 실행(.Do)
			b = &blockchain{} // 2. blockchain 인스턴스 생성
			b.AddBlock("Genesis")
		})
	}
	return b // 3. b 반환
}

func (b *blockchain) AllBlocks() []*block {
	return GetBlockchain().blocks
}
