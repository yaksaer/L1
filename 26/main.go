package main

import "fmt"

// checkUniq check symbols in string for uniqueness using map and ASCII table
func checkUniq(str string) bool {
	//map for checking symbols
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
	//create string
	str := "abcdefgAtiok"
	fmt.Println("String with all uniq characters")
	//check it and print result
	fmt.Println("Result:", checkUniq(str))
	//create string
	str = "abCdeFgijklmoPqrstbA"
	fmt.Println("String with not uniq characters")
	//check it and print result
	fmt.Println("Result:", checkUniq(str))
}
