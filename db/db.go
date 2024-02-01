package db

import (
	"github.com/boltdb/bolt"
	"github.com/jhin93/Golang_nomadCoin/utils"
)

// blockchain.go 처럼 singleton 패턴을 사용

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open("blockchain.db", 0600, nil) // https://github.com/boltdb/bolt?tab=readme-ov-file#opening-a-database
		utils.HandleErr(err)
		db = dbPointer
	}
	return db
}
