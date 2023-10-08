package main

import (
	"fmt"
	"sync"
	"time"
)

// creacional pattern
type Database struct{}

var db *Database
var lock sync.Mutex

// create an instance
func (Database) CreateSingleConnection() {
	fmt.Println("Create a DB instance for singleton")
	time.Sleep(2 * time.Second)
	fmt.Println("Creation Done")

}

// get database instance
func getDataBaseInstance() *Database {
	// lock until the instance is created
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB alreasy created")
	}
	return db
}

func main() {
	// use waitgroups for concurrency
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDataBaseInstance()
		}()
	}
	wg.Wait()
}
