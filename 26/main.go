package main

import "fmt"

func checkUniq(str string) bool {
	buf := make(map[uint]int, len(str))
	for i := range str {
		if buf[uint(str[i])] == 1 {
			return false
		}
		buf[uint(str[i])] = 1
	}
	return true
}

func main() {
	str := "abcdefgAtiok"
	fmt.Println("String with all uniq characters")
	fmt.Println("Result:", checkUniq(str))
	str = "abCdeFgijklmoPqrstbA"
	fmt.Println("String with not uniq characters")
	fmt.Println("Result:", checkUniq(str))
}
