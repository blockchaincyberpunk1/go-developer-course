package main

import (
	"fmt"
	"sync"
)

// Function to calculate the sum of a portion of the array.
func calculatePartialSum(arr []int, startIdx int, endIdx int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done() // Notify that this goroutine is done when it exits.

	partialSum := 0
	for i := startIdx; i < endIdx; i++ {
		partialSum += arr[i]
	}

	// Send the partial sum to the result channel.
	resultChan <- partialSum
}

func main() {
	// Create a large array of numbers.
	arr := make([]int, 1000000)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	// Define the number of goroutines to use.
	numGoroutines := 4

	// Create a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// Create a channel to receive partial sums from goroutines.
	resultChan := make(chan int, numGoroutines)

	// Split the array into equal portions for each goroutine.
	partitionSize := len(arr) / numGoroutines

	// Launch goroutines to calculate partial sums concurrently.
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) // Increment the WaitGroup counter.

		startIdx := i * partitionSize
		endIdx := (i + 1) * partitionSize

		// For the last goroutine, adjust the end index to cover the remaining elements.
		if i == numGoroutines-1 {
			endIdx = len(arr)
		}

		go calculatePartialSum(arr, startIdx, endIdx, &wg, resultChan)
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	// Close the result channel since no more values will be sent.
	close(resultChan)

	// Collect and sum the partial sums from the result channel.
	totalSum := 0
	for partialSum := range resultChan {
		totalSum += partialSum
	}

	// Print the final sum.
	fmt.Printf("Total sum: %d\n", totalSum)
}
