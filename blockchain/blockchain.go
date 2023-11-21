package blockchain

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

var b *blockchain

func GetBlockchain() *blockchain { // 변수 b와 동일한 타입인 blockchain의 pointer 반환
	// 1. b 변수 초기화 여부 확인
	if b == nil {
		b = &blockchain{} // 2. blockchain 인스턴스 생성
	}
	return b // 3. b 반환
}

// 누군가 이 blockchain을 처음으로 요청하면 nil을 반환하는 일 없이 blockchain을 먼저 초기화 한 뒤에 반환함
// 이미 실행된 blockchain을 누군가 초기화하려고 해도 이미 실행됐던 GetBlockchain 함수로 인해 이미 nil이 아닌 b가 반환된다
