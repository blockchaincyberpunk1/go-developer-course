package main

import (
	"fmt"
	"sync"
)

// Function that simulates a time-consuming task.
func performTask(taskID int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify that this task is done when the function exits.

	// Simulate some work
	fmt.Printf("Task %d: Started\n", taskID)
	for i := 0; i < 3; i++ {
		fmt.Printf("Task %d: Working...\n", taskID)
	}
	fmt.Printf("Task %d: Completed\n", taskID)
}

func main() {
	// Create a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// Introduce the concept of goroutines.
	fmt.Println("Introduction to Goroutines:")
	fmt.Println("Goroutines are lightweight threads of execution.")

	// Launch multiple goroutines concurrently.
	for i := 1; i <= 3; i++ {
		wg.Add(1)              // Increment the WaitGroup counter.
		go performTask(i, &wg) // Start a goroutine to perform a task.
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	// Introduce the concept of channels.
	fmt.Println("\nIntroduction to Channels:")
	fmt.Println("Channels are used for communication between goroutines.")

	// Create a channel for communication.
	ch := make(chan string)

	// Start a goroutine to send a message to the channel.
	go func() {
		ch <- "Hello, from a channel!"
	}()

	// Receive the message from the channel.
	msg := <-ch
	fmt.Println(msg)

	// Discussion:
	fmt.Println("\nDiscussion:")
	fmt.Println("Concurrency allows for efficient use of CPU resources.")
	fmt.Println("Benefits include improved program performance.")
	fmt.Println("Challenges include data races and synchronization issues.")
	fmt.Println("Concurrency is crucial in scenarios like web servers and data processing.")
}
