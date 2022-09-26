package main

import "fmt"

//createSubset check values from str and create its subset of uniq values
func createSubset(str []string) []string {
	//create map for remember all uniq values
	checkSet := make(map[string]int)
	//create result string
	var result []string
	//take all values from str and save them as key in map, set key's values as '1'
	for _, v := range str {
		checkSet[v] = 1
	}
	//take all map's keys and create new array of uniq values
	for key, _ := range checkSet {
		result = append(result, key)
	}
	return result
}

func main() {
	//create array of string
	str := []string{"cat", "cat", "dog", "cat", "tree", "tree", "dog"}
	fmt.Println(createSubset(str))
}
