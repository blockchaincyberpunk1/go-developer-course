package main

import "fmt"

// Task a: Function to calculate the square of a number
func square(x int) int {
	return x * x
}

// Task b: Function to determine if a number is even or odd
func evenOrOdd(num int) string {
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// Task c: Function to calculate the sum of two numbers
func sum(a, b int) int {
	return a + b
}

func main() {
	// Task a: Calculate the square of a number
	num := 5
	squared := square(num)
	fmt.Printf("Task a: Square of %d is %d\n", num, squared)

	// Task b: Determine if a number is even or odd
	num2 := 8
	result := evenOrOdd(num2)
	fmt.Printf("Task b: %d is %s\n", num2, result)

	// Task c: Calculate the sum of two numbers
	x := 10
	y := 20
	total := sum(x, y)
	fmt.Printf("Task c: Sum of %d and %d is %d\n", x, y, total)
}
