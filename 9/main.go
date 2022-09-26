package main

import (
	"fmt"
	"sync"
)

// Channels with all needed channels and variable to wait all goroutines
type Channels struct {
	number chan int
	result chan int
	wg     sync.WaitGroup
}

// New create new Channels object
func New() Channels {
	return Channels{
		number: make(chan int),
		result: make(chan int),
		wg:     sync.WaitGroup{},
	}
}

// writeInNumber take numbers from array and write them into number channel
func writeInNumber(numbers []int, channels *Channels) {
	defer channels.wg.Done()
	defer close(channels.number)
	for _, v := range numbers {
		channels.number <- v
	}
}

//multiply read numbers from number channel, multiply them by 2 and write them into result channel
func multiply(channels *Channels) {
	defer channels.wg.Done()
	defer close(channels.result)
	for v := range channels.number {
		channels.result <- v
	}
}
//writeInStdout read numbers from result channel and write them in STDOUT
func writeInStdout(channels *Channels) {
	defer channels.wg.Done()
	for v := range channels.result {
		fmt.Println(v)
	}
}

func main() {
	//create new channels object
	channels := New()
	//create and fill array of int
	numbers := make([]int, 20)
	for i := 0; i < 20; i++ {
		numbers[i] = i
	}
	//start all needed goroutines
	channels.wg.Add(3)
	go writeInNumber(numbers, &channels)
	go multiply(&channels)
	go writeInStdout(&channels)
	//wait all goroutines
	channels.wg.Wait()
}
