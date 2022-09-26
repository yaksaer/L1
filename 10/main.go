package main

import "fmt"

func main() {
	//create array of temperatures
	temperatures := []float64{-25.4, -27.0, 13.0, -40.1, -41.9, 31, 4, 6, -11, 19.0, 15.5, 24.5, -21.0, 32.5}
	//create map of results
	result := make(map[int][]float64)
	for _, v := range temperatures {
		//check which group the value belongs to
		key := int(v / 10)
		//put value into result map
		result[key*10] = append(result[key*10], v)
	}
	//print result
	fmt.Println(result)
}
