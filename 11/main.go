package main

import "fmt"

// findIntersection write all values from str1 into the map as key and then check str2 on having the same values
func findIntersection(str1 []int, str2 []int) []int {
	//create map to remember all values from str1
	checkInter := make(map[int]int)
	//result array
	var result []int
	for _, v := range str1 {
		//write value from str1 as key and '1' like a flag
		checkInter[v] = 1
	}
	for _, v := range str2 {
		//check str2's values, if we have '1' for this key, its mean the value is repeats
		if checkInter[v] == 1 {
			//add reapeats value into result array
			result = append(result, v)
		}
	}
	return result
}

func main() {
	//create two arrays
	str1 := []int{1, -9, 6, 8, 14, 43, 98, -5, 7, 2, -93, 35, 12, 20}
	str2 := []int{1, 76, 412, 8, -84, 34, 16, 6, 7, 0, 84, 93, 10, 20}
	fmt.Println(findIntersection(str1, str2))
}
