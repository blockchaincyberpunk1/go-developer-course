package main

import (
	"fmt"
	"strings"
)

// Function 1 - Input and Validation
func getUserInput() string {
	var name string
	for {
		// Prompt the user to enter their name with a minimum length requirement.
		fmt.Print("Enter your name (at least 2 characters): ")
		// Read the user's input from the console.
		fmt.Scanln(&name)
		// Remove leading and trailing spaces from the input.
		name = strings.TrimSpace(name)
		if len(name) >= 2 {
			// If the input has at least 2 characters, return it.
			return name
		}
		// Display an error message if the input is too short.
		fmt.Println("Name must be at least 2 characters long. Try again.")
	}
}

// Function 2 - Calculation
func calculateRectangleArea(length, width float64) (float64, error) {
	if length <= 0 || width <= 0 {
		// Check if the provided dimensions are non-positive.
		return 0, fmt.Errorf("length and width must be positive numbers")
	}
	// Calculate the area of the rectangle.
	return length * width, nil
}

// Function 3 - Decision-Making
func isEvenOrOdd(number int) string {
	if number%2 == 0 {
		// Check if the number is even.
		return "even"
	}
	// If not even, it's odd.
	return "odd"
}

// Function 4 - Display
func displayWelcomeMessage(name string, result string) {
	// Display a welcome message with the user's name and the result.
	fmt.Printf("Welcome, %s!\nYour chosen number is %s.\n", name, result)
}

func main() {
	// Function 1 - Input and Validation
	name := getUserInput()

	// Function 2 - Calculation
	length := 5.0
	width := 4.0
	// Calculate the area of a rectangle and handle errors.
	_, err := calculateRectangleArea(length, width)
	if err != nil {
		// If there is an error, display it and exit the program.
		fmt.Printf("Error: %s\n", err)
		return
	}

	// Function 3 - Decision-Making
	number := 7
	result := isEvenOrOdd(number)

	// Function 4 - Display
	// Display a welcome message with the user's name and the result.
	displayWelcomeMessage(name, result)
}
