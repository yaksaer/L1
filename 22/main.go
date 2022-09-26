package main

import (
	"fmt"
	"math/big"
)

func main() {
	//we can work with big integers by big package
	a, b := new(big.Float), new(big.Float)
	a.SetString("1923882319382173871387238173812931867523")
	b.SetString("9876543211234567899876543211234567899876")
	result := new(big.Float)
	fmt.Println("Calculation with big package\na + b =", result.Add(a, b))
	fmt.Println("a - b =", result.Sub(a, b))
	fmt.Println("a * b =", result.Mul(a, b))
	fmt.Println("a / b =", result.Quo(a, b))
	//also untyped int constant can contain big integer values, bigger than 2^20 but smaller than big int package
	const c, d = 192388231, 9876543211
	fmt.Println("\nCalculation with untyped constant\nc + d =", c+d)
	fmt.Println("c - d =", c-d)
	fmt.Println("c * d =", c*d)
	fmt.Println("c / d =", c/d)
}
