package main

import "fmt"

func main() {
	// List of numbers
	numbers := []int{1, 2, 3, 4, 5}

	// Operation 1: Calculate the sum of all elements
	sum := 0
	// Iterate through the numbers slice using a for-each loop.
	for _, num := range numbers {
		// Add the current number to the running total (sum).
		sum += num
	}
	// Print the result of Operation 1.
	fmt.Printf("Operation 1: Sum of all elements: %d\n", sum)

	// Operation 2: Find the maximum value
	max := numbers[0]
	// Iterate through the numbers slice again.
	for _, num := range numbers {
		// Compare the current number with the current maximum.
		if num > max {
			// If the current number is greater, update the maximum.
			max = num
		}
	}
	// Print the result of Operation 2.
	fmt.Printf("Operation 2: Maximum value: %d\n", max)

	// Operation 3: Count the even numbers
	evenCount := 0
	// Iterate through the numbers slice one more time.
	for _, num := range numbers {
		// Check if the current number is even (divisible by 2).
		if num%2 == 0 {
			// If it's even, increment the evenCount.
			evenCount++
		}
	}
	// Print the result of Operation 3.
	fmt.Printf("Operation 3: Count of even numbers: %d\n", evenCount)

	// Operation 4: Reverse the order of elements
	// Create a new slice to store the reversed numbers with the same length as the original slice.
	reversedNumbers := make([]int, len(numbers))
	// Use two index variables, i and j, to iterate through the original and reversed slices.
	for i, j := 0, len(numbers)-1; i < len(numbers); i, j = i+1, j-1 {
		// Assign the current element from the original slice to the reversed slice in reverse order.
		reversedNumbers[i] = numbers[j]
	}
	// Print the result of Operation 4.
	fmt.Printf("Operation 4: Reversed list: %v\n", reversedNumbers)
}
