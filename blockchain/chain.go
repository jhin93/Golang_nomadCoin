package blockchain

import (
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once // sync 패키지의 type 'Once'

func (b *blockchain) AddBlock(data string) {
	createBlock(data, b.NewestHash, b.Height)
}

func Blockchain() *blockchain { // 변수 b와 동일한 타입인 blockchain의 pointer 반환
	// 1. b 변수 초기화 여부 확인
	if b == nil {
		once.Do(func() { // 내부의 로직을 오직 한번만 실행(.Do)
			b = &blockchain{"", 0} // 2. blockchain 인스턴스 생성. 첫 블록 생성때는 NewestHash에는 아무것도 없고(""), Height는 0이다.
			b.AddBlock("Genesis")
		})
	}
	return b // 3. b 반환
}
