package main

import (
	"context"
	"fmt"
	"time"
)

type channels struct {
	message chan string
	stop    chan bool
}

// initStruct initialization of all needed channels
func initStruct() channels {
	channels := channels{}
	channels.message = make(chan string)
	channels.stop = make(chan bool)
	return channels
}

// work execution of goroutine
func work(channels channels) {
	for {
		buf, ok := <-channels.message
		//stop if channel is closed
		if !ok {
			fmt.Println("Goroutine has been stopped")
			return
		}
		fmt.Println(buf)
	}
}

func work1(channels channels) {
	for {
		select {
		//case if we receive message
		case buf := <-channels.message:
			fmt.Println(buf)
		//case if the stop-message has been received
		case <-channels.stop:
			fmt.Println("Goroutine has been stopped")
			return
		}
	}
}

func work2(channels channels, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine has been stopped")
			return
		case buf := <-channels.message:
			fmt.Println(buf)
		}
	}
}

func main() {
	channels := initStruct()

	//Stop goroutine by closing channel
	go work(channels)
	channels.message <- "Hello from goroutine number 1"
	close(channels.message)
	time.Sleep(2 * time.Second)

	//Stop by message from channel
	channels.message = make(chan string)
	go work1(channels)
	channels.message <- "Hello from goroutine number 2"
	channels.stop <- true
	time.Sleep(2 * time.Second)

	//Stop by context cancel
	ctx, cancel := context.WithCancel(context.Background())
	go work2(channels, ctx)
	channels.message <- "Hello from goroutine number 3"
	cancel()
	time.Sleep(2 * time.Second)
}
