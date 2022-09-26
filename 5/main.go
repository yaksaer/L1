package main

import (
	"fmt"
	"strconv"
	"time"
)

// Struct with all needed channels
type channels struct {
	message chan string
	stop    chan bool
}

// work execution of goroutine
func work(channels channels) {
	for {
		select {
		//case if we receive message
		case buf := <-channels.message:
			fmt.Println("Message has been received:", buf)
			time.Sleep(1000000)
		//case if time has passed, so the stop-message has been received
		case <-channels.stop:
			fmt.Println("Worker number finished his work")
			return
		}
	}
}

// initStruct initialization of all needed channels
func initStruct() channels {
	channels := channels{}
	channels.message = make(chan string)
	channels.stop = make(chan bool)
	return channels
}

func main() {
	fmt.Println("Please insert after how many second finish the program:")
	workTime := 0
	//Read number of seconds from STDIN
	_, err := fmt.Scan(&workTime)
	if err != nil {
		return
	}
	//Initialization of all channels
	channels := initStruct()
	defer close(channels.stop)
	defer close(channels.message)
	//Save current time
	start := time.Now()
	//Start goroutine
	go work(channels)
	//Send message to goroutine and stop them if time has passed
	for {
		channels.message <- "Hello, how are you?"
		passedTime := time.Now().Second() - start.Second()
		if passedTime >= workTime {
			channels.stop <- true
			fmt.Println("\033[32m", "Program has been stopped", "\033[0m")
			fmt.Println("\033[36m", "Seconds have passed since the start", strconv.Itoa(passedTime)+"s", "\033[0m")
			return
		}
	}
}
