package main

import "fmt"

func main() {
	str := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	i := 4
	fmt.Println("Start slice:", str, "\ni =", i)
	//copy last element in i index
	str[i] = str[len(str)-1]
	//delete last element
	str[len(str)-1] = ""
	//truncate slice
	str = str[:len(str)-1]
	fmt.Println("Slice after delete:", str)
	str = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	i = 2
	fmt.Println("Start slice:", str, "\ni =", i)
	//shift left by one index
	copy(str[i:], str[i+1:])
	//delete last element
	str[len(str)-1] = ""
	//truncate slice
	str = str[:len(str)-1]
	fmt.Println("Slice after delete:", str)
}
