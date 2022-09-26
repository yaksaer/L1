package main

import "fmt"

type nums struct {
	num       int64
	bitNum    int
	targetNum int
}

func changeBit(nums nums) int64 {
	var mask byte
	mask = 1 << nums.bitNum
	return nums.num ^ int64(mask)
}

func readSTDIN() (nums, error) {
	nums := nums{}
	fmt.Println("Insert int64 variable:")
	//Read numbers from STDIN
	_, err := fmt.Scan(&nums.num)
	if err != nil {
		return nums, err
	}
	fmt.Println("Insert the number of bit you want to change:")
	_, err = fmt.Scan(&nums.bitNum)
	if err != nil {
		return nums, err
	}
	fmt.Println("Insert the number to which you want to change this bit ('1' or '0'):")
	_, err = fmt.Scan(&nums.targetNum)
	if err != nil {
		return nums, err
	}
	return nums, err
}

func main() {
	nums, err := readSTDIN()
	if err != nil {
		return
	}
	if nums.targetNum != 0 && nums.targetNum != 1 {
		fmt.Println("Wrong bit's desired value. Must be 1 or 0!")
		return
	}
	res := changeBit(nums)
	fmt.Printf("Your value:%b\nNew value: %b\n", nums.num, res)
}
