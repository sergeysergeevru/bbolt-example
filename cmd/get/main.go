package main

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)


func main() {
	db, err := bolt.Open("test.db", 0666,nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("the db has been opened")

	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))

		fmt.Printf("%s\n", b.Get([]byte("answer")))
		return nil
	})
	if err != nil {
		panic(err)
	}
}


