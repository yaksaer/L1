package main

import "fmt"

// quickSort sort array using quick sort algorithm
func quickSort(arr []int, start int, end int) {
	//if array size is 1 or 0 no sort needed
	if (end - start) <= 1 {
		return
	}
	//define our pivot value
	pivot := arr[end]
	//define splitIndex which will contain index value of position in the array that splits it into two new sub-arrays
	splitNum := start
	/*
		Then we iterate over our subarray defined by function parameters,
		and we check if value at this index is lower than pivot value,
		if so we move it to the beginning of an array, switching positions with element at splitIndex.
		Then we increase our splitIndex value by 1 to indicate that before this index there are values
		less-than pivot and after this index there are only values greater-than pivot.
		Then we move our pivot value next to last value that was lower than it which is at splitIndex
	*/
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			tmp := arr[splitNum]
			arr[splitNum] = arr[i]
			arr[i] = tmp
			splitNum++
		}
	}
	/*
		When function finishes its execution, subarray contains sorted values.
		Finally, we want to recursively call our sort function for created sub-arrays.
		So it applies the same process, finding all lower-greater values
		and defining sub-arrays till every element in our main array is sorted.
	*/
	arr[end] = arr[splitNum]
	arr[splitNum] = pivot
	quickSort(arr, start, splitNum-1)
	quickSort(arr, splitNum+1, end)
}

func main() {
	arr := []int{9, 34, 1, 0, -5, 12, 43, 67, -50, 40, 5, 3, 91, -12, -1, 98, 65, 51}
	fmt.Println("Before sort:\n", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("After sort:\n", arr)
}
