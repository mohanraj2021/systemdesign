package main

import (
	"database/sql"
	"fmt"
	"sync"
)

var once sync.Once
var dbInstance *singleDbInstance

// var mu sync.Mutex

type singleDbInstance struct {
	Db *sql.DB
}

func GetDbInstance() *singleDbInstance {
	// logic to create only one instance of DB connection
	once.Do(func() {
		dbInstance = &singleDbInstance{Db: nil}
	})
	return dbInstance
}

func main() {
	for i := 0; i < 10000; i++ {
		go func() {
			db := GetDbInstance()
			db2 := GetDbInstance()
			if db != db2 {
				fmt.Println("not same")
			}
		}()
	}

}
