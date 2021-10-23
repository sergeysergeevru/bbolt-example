package main

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)


func main() {
	db, err := bolt.Open("test.db", 0666, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("the db has been opened")

	err = db.Update(func(tx *bolt.Tx) error {
		bucketName := []byte("MyBucket")
		b := tx.Bucket(bucketName)
		if b == nil {
			b, err = tx.CreateBucket(bucketName)
			if err != nil {
				return err
			}
		}

		return b.Put([]byte("answer"), []byte("some test data"))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("the data updated")
}


