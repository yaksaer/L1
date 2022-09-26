package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	//wait a specified amount of time
	<-time.After(d * time.Second)
}

func main() {
	fmt.Println("Going to sleep for 3 seconds")
	sleep(3)
	fmt.Println("Time to get up")
}
