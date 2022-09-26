package main

import (
	"fmt"
	"strconv"
)

type arg struct {
	str []string
	res int
	ch  chan int
}

// pow calculates square power of the given number and write it into channel
func pow(num string, ch chan int) {
	n, _ := strconv.Atoi(num)
	n *= n
	ch <- n
}

func main() {
	//Initialisation data and channel
	data := arg{}
	data.ch = make(chan int)
	defer close(data.ch)
	//Uncomment this lines and comment line with data.str to use STDIN instead of the given parameters
	/*
		fmt.Println("Введите ряд чисел разделяя их пробелом:")
		scanner := bufio.NewReader(os.Stdin)
		buf, _ := scanner.ReadString('\n')
		data.str = strings.Split(buf[:len(buf)-1], " ")
	*/

	data.str = []string{"2", "4", "6", "8", "10"}
	res := 0
	for _, v := range data.str {
		//Start goroutines
		go pow(v, data.ch)
	}
	for _, _ = range data.str {
		//Waiting until goroutine write information in channel, read it and add to result value
		res += <-data.ch
	}
	//Print result value
	fmt.Println("Сумма квадратов заданных чисел =", res)
}
