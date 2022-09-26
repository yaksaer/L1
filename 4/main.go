package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Struct with all needed channels
type channels struct {
	message     chan string
	catchSig    chan os.Signal
	stopWorkers chan bool
}

// work execution of goroutine
func work(channels channels, n int) {
	//Initialization color variable for different color output for workers
	color := ""
	if n%2 == 0 {
		color = "\033[35m"
	} else {
		color = "\033[36m"
	}
	//Read message from channel and stop worker if message has been found in stopWorkers channel
	for {
		select {
		case buf := <-channels.message:
			fmt.Println(color, "Worker number", n, "is working with", buf, "\033[0m")
			//Uncomment this line for slower output
			//time.Sleep(1000000)
		case <-channels.stopWorkers:
			fmt.Println(color, "Worker number", n, "finished his work", "\033[0m")
			return
		}
	}
}

// initStruct initialization of all needed channels
func initStruct(workersNum int) channels {
	channels := channels{}
	channels.message = make(chan string, workersNum)
	channels.catchSig = make(chan os.Signal)
	channels.stopWorkers = make(chan bool)
	return channels
}

func main() {
	fmt.Println("Insert number of workers:")
	//Read number of workers from STDIN
	workersNum := 0
	_, err := fmt.Scan(&workersNum)
	if err != nil {
		return
	}
	//Initialization of all channels
	channels := initStruct(workersNum)
	defer close(channels.stopWorkers)
	defer close(channels.catchSig)
	defer close(channels.message)
	//Notify about signal
	signal.Notify(channels.catchSig, syscall.SIGINT, syscall.SIGTERM)
	//Start all goroutines
	for i := 0; i < workersNum; i++ {
		go work(channels, i+1)
	}
	//Send message to goroutines and stop them if signal has come
	for {
		select {
		case channels.message <- "papers":
		case <-channels.catchSig:
			for i := 0; i < workersNum; i++ {
				channels.stopWorkers <- true
			}
			fmt.Println("\033[32m", "All workers have finished their work", "\033[0m")
			return
		}
	}
}
