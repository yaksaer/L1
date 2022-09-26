package main

import "fmt"

// reverse the string
func reverse(str string) string {
	//change str type to []rune
	buf := []rune(str)
	//run from start and end of the array and swap letters
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	//return result
	return string(buf)
}

func main() {
	str := "водогрязеторфопарафинолечение"
	fmt.Printf("Исходная строка:\n%s\nПеревёрнутая строка:\n%s\n", str, reverse(str))
	str = "antiskepticism"
	fmt.Printf("Исходная строка:\n%s\nПеревёрнутая строка:\n%s\n", str, reverse(str))
}
