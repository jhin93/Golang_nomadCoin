package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []*block // blocks는 pointer의 slice
}

var b *blockchain
var once sync.Once // sync 패키지의 type 'Once'

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash) // 16진수
}

func getLastHash() string { // 마지막 블록 해쉬 반환
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].hash
}

func createBlock(data string) *block { // block 타입의 pointer 반환
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock // 반환값이 메모리 연산자인 이유는 createBlock 함수가 사용될 blockchain구조체의 내부 요소 blocks가 포인터 변수 슬라이스이기 때문
}

func GetBlockchain() *blockchain { // 변수 b와 동일한 타입인 blockchain의 pointer 반환
	// 1. b 변수 초기화 여부 확인
	if b == nil {
		once.Do(func() { // 내부의 로직을 오직 한번만 실행(.Do)
			b = &blockchain{} // 2. blockchain 인스턴스 생성
			b.blocks = append(b.blocks, createBlock("Genesis Block"))
		})
	}
	return b // 3. b 반환
}

// 누군가 이 blockchain을 처음으로 요청하면 nil을 반환하는 일 없이 blockchain을 먼저 초기화 한 뒤에 반환함
// 이미 실행된 blockchain을 누군가 초기화하려고 해도 이미 실행됐던 GetBlockchain 함수로 인해 이미 nil이 아닌 b가 반환된다
