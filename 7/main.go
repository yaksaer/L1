package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Storage struct with map, mutex to protect it and variable to wait all goroutines
type Storage struct {
	sync.Mutex
	data map[int]string
	wg   sync.WaitGroup
}

// New create new Storage object
func New() *Storage {
	return &Storage{
		data: make(map[int]string),
	}
}

// set add value to the map
func (storage *Storage) set(key int, value string) {
	storage.data[key] = value
}

// fillMap add value in map concurrently
func fillMap(storage *Storage, key int) {
	//defer done for end goroutine message
	defer storage.wg.Done()
	//lock mutex for add value into the map
	storage.Lock()
	//defer to unlock mutex after adding
	defer storage.Unlock()
	//add value into the map
	storage.set(key, "value"+strconv.Itoa(key))
}

func main() {
	//create new storage variable
	storage := New()
	//cycle with goroutine start and adding values into the map
	for i := 0; i < 10; i++ {
		storage.wg.Add(1)
		go fillMap(storage, i)
	}
	//wait all goroutines
	storage.wg.Wait()
	fmt.Println(storage.data)
}
