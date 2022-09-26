package main

import "fmt"

/*First version is using []byte instead of string type, because when we change string, Golang create new one, and it's
very expensive with big strings*/

// createHugeString create []byte with strLen elements
func createHugeString(strLen int) []byte {
	buf := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		buf[i] = 'J'
	}
	return buf
}

// someFunc return first 100 elements of string
func someFunc() []byte {
	return createHugeString(1 << 10)[:100]
}

/*Second version is using string type*/

// createHugeString create string with strLen elements
func createHugeString2(strLen int) string {
	var buf string
	for i := 0; i < strLen; i++ {
		buf += "J"
	}
	return buf
}

// someFunc return first 100 elements of string
func someFunc2() string {
	return createHugeString2(1 << 10)[:100]
}

func main() {
	fmt.Printf("Variant 1:\n%s\n", someFunc())
	fmt.Printf("Variant 2:\n%s\n", someFunc2())
}
