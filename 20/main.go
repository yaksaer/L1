package main

import "fmt"

// reverseStrings revers words in array of strings
func reverseStrings(str []string) []string {
	//create result variable
	result := make([]string, len(str))
	//run from start and end of the array and swap words
	for j, i := len(str)-1, 0; j >= 0; j, i = j-1, i+1 {
		result[i] = str[j]
	}
	//return result
	return result
}

func main() {
	str := []string{"Всем", "привет", "меня", "зовут", "Боб"}
	fmt.Printf("Исходная строка:\n%s\nПеревёрнутая строка:\n%s\n", str, reverseStrings(str))
	str = []string{"Hello", "everyone", "my", "name", "is", "Bob"}
	fmt.Printf("Исходная строка:\n%s\nПеревёрнутая строка:\n%s\n", str, reverseStrings(str))
}
