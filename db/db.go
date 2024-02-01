// blockchain.go 처럼 singleton 패턴을 사용
package db

import (
	"github.com/boltdb/bolt"
	"github.com/jhin93/Golang_nomadCoin/utils"
)

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB // export 되지 않는 초기 변수.

func DB() *bolt.DB {
	if db == nil {
		// 1. db가 존재하지 않으면 dbPointer에 지정
		dbPointer, err := bolt.Open(dbName, 0600, nil) // https://github.com/boltdb/bolt?tab=readme-ov-file#opening-a-database
		db = dbPointer
		// 2. 에러 처리
		utils.HandleErr(err)
		// 3. Bucket이 존재하는지 확인. 존재하지 않으면 생성(dataBucket, blocksBucket)시켜주는 transaction. Bucket은 SQL의 table과 같은 것
		err = db.Update(func(t *bolt.Tx) error { // https://github.com/boltdb/bolt?tab=readme-ov-file#read-write-transactions
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket)) // https://github.com/boltdb/bolt?tab=readme-ov-file#using-buckets
			utils.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket)) // err가 2줄 위에 이미 선언(err := t.Create...) 되었기에 그냥 할당만 함.
			return err                                               // 1줄 위 마지막 에러를 리턴
		})
		utils.HandleErr(err)
	}
	return db
}
