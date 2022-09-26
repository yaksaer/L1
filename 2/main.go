package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type arg struct {
	str []string
	wg  sync.WaitGroup
}

func pow(num string, data *arg) {
	defer data.wg.Done()
	n, _ := strconv.Atoi(num)
	n *= n
	fmt.Println(num, "*", num, "=", n)
}

func main() {
	data := arg{}
	fmt.Println("Введите ряд чисел разделяя их пробелом:")
	scanner := bufio.NewReader(os.Stdin)
	buf, _ := scanner.ReadString('\n')
	data.str = strings.Split(buf[:len(buf)-1], " ")
	for _, v := range data.str {
		data.wg.Add(1)
		go pow(v, &data)
	}
	data.wg.Wait()
}
