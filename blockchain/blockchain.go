package blockchain

import (
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

type blockchain struct {
}

var b *blockchain
var once sync.Once // sync 패키지의 type 'Once'

func (b *blockchain) AddBlock(data string) {

}

func Blockchain() *blockchain { // 변수 b와 동일한 타입인 blockchain의 pointer 반환
	// 1. b 변수 초기화 여부 확인
	if b == nil {
		once.Do(func() { // 내부의 로직을 오직 한번만 실행(.Do)
			b = &blockchain{} // 2. blockchain 인스턴스 생성
			b.AddBlock("Genesis")
		})
	}
	return b // 3. b 반환
}
