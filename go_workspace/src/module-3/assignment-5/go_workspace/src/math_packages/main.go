package main

import (
	"fmt"
	"math_packages/mathlib"  // Import the mathlib package with alias "mlib"
	"math_packages/mathutil" // Import the mathutil package without alias
)

func main() {
	a, b := 10, 5

	// Use functions from both packages
	sum := mathlib.Add(a, b)
	product := mathutil.Multiply(a, b)

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Product: %d\n", product)
}
