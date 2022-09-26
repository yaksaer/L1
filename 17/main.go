package main

import (
	"fmt"
	"sort"
)

// binarySearch looking for x value in the array, if x has been found return true, else return false
func binarySearch(arr []int, x int) bool {
	//create limit values as start and end of the array
	low, high := 0, len(arr)-1
	for low < high {
		//evaluate middle of the array
		middle := int(uint(low+high) / 2)
		if x > arr[middle] {
			low = middle + 1 //if x is bigger than middle, set low variable higher than middle
		} else {
			high = middle //if x is smaller or equal the middle, set high variable equal to middle
		}
	}
	//if we didn't find x value return false
	if arr[low] != x || low == len(arr) {
		return false
	}
	return true
}

func main() {
	arr := []int{9, 34, 1, 0, -5, 12, 43, 67, -50, 40, 5, 3, 91, -12, -1, 98, 65, 51}
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(binarySearch(arr, 9))
}
