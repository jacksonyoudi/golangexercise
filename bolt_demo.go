package main

import (
	"github.com/boltdb/bolt"
	"os"
	"fmt"
)

const MYDB  = "my.db"
const  MYBUCKET = "mybucket"


func main()  {
	db, err := bolt.Open(MYDB, 777, nil)
	if err != nil {
		panic("open db failure")
		os.Exit(1)
	}

	/*db.Update(func(tx *bolt.Tx) error {
		bk, err := tx.CreateBucket([]byte(MYBUCKET))
		if err != nil {
			panic("create bucket failure")
			os.Exit(1)
		}
		bk.Put([]byte("hello"), []byte("world"))
		return nil
	})*/

	db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(MYBUCKET))
		v := bk.Get([]byte("hello"))
		fmt.Println(string(v))
		return nil
	})

	db.Close()
}
