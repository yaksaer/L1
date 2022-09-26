package main

import "fmt"

func main() {
	//create variables and swap their value by multi-value assignments
	fmt.Println("Swap two string variables")
	a, b := "one", "two"
	fmt.Printf("Start values:a = %s ; b = %s\n", a, b)
	a, b = b, a
	fmt.Printf("After swap:  a = %s ; b = %s\n\n", a, b)
	fmt.Println("Swap two int variables")
	c, d := 1, 2
	fmt.Printf("Start values:c = %d ; d = %d\n", c, d)
	c, d = d, c
	fmt.Printf("After swap:  c = %d ; d = %d\n\n", c, d)
	fmt.Println("Swap elements in array of int")
	str := []int{10, 20, 30, 40, 50}
	fmt.Println("Start values:", str)
	str[0], str[4] = str[4], str[0]
	fmt.Println("After swap:  ", str, "\n")
	fmt.Println("Swap elements in array of string")
	str2 := []string{"Hello", "my", "name", "is", "Bob"}
	fmt.Println("Start values:", str2)
	str2[0], str2[4] = str2[4], str2[0]
	fmt.Println("After swap:  ", str2)
}
