package main

import (
	"assignment-2/custommath" // Import the custommath package
	"fmt"
	"log"
)

func main() {
	// Using functions from the custommath package
	sum := custommath.Add(10, 5)
	difference := custommath.Subtract(15, 7)

	result, err := custommath.Divide(20, 4)
	if err != nil {
		log.Fatal(err)
	}

	// Displaying results
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Division Result:", result)
}
