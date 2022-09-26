package main

import (
	"fmt"
	"sync"
)

// Counter struct with mutex and counter
type Counter struct {
	sync.Mutex
	i  int
	wg sync.WaitGroup
}

// New create new Counter object
func New() *Counter {
	return &Counter{
		i: 0,
	}
}

// count increase counter
func (storage *Counter) count() {
	storage.Lock()
	storage.i++
	storage.Unlock()
}

func main() {
	//create new counter object
	counter := New()
	//start all goroutines in cycle
	for i := 0; i < 200; i++ {
		counter.wg.Add(1)
		go func() {
			//increase counter
			counter.count()
			counter.wg.Done()
		}()
	}
	//wait all goroutines
	counter.wg.Wait()
	//print result
	fmt.Println(counter.i)
}
