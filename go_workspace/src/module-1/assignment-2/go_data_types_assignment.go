package main

import (
	"fmt"
)

func main() {
	// Declare variables of different data types
	age := 25
	temperature := 98.6
	name := "Alice"
	isStudent := true

	// Print the values and their respective data types
	fmt.Printf("Age: %v (%T)\n", age, age)
	fmt.Printf("Temperature: %v (%T)\n", temperature, temperature)
	fmt.Printf("Name: %v (%T)\n", name, name)
	fmt.Printf("Is Student: %v (%T)\n", isStudent, isStudent)

	// Importance of Data Types
	fmt.Println("\nImportance of Data Types:")
	fmt.Println("Choosing the right data type is crucial for:")
	fmt.Println("- Memory Usage: Inefficient data types can waste memory.")
	fmt.Println("- Type Safety: Ensuring that operations are performed on compatible data types.")
	fmt.Println("- Program Efficiency: Using appropriate data types can lead to faster code execution.")
	fmt.Println("- Code Clarity: Selecting meaningful data types makes the code more understandable.")
	fmt.Println("- Error Prevention: Helps catch type-related errors during compilation.")

	// Scenarios
	fmt.Println("\nScenarios:")

	// Scenario 1
	fmt.Println("Scenario 1: Shopping Cart")
	fmt.Println("For item prices and total price, you should use float64 data type.")

	// Scenario 2
	fmt.Println("\nScenario 2: Weather Forecasting")
	fmt.Println("For temperature readings in Celsius, use float64.")
	fmt.Println("For precipitation forecasts, use string (e.g., rainy, sunny).")

	// Scenario 3
	fmt.Println("\nScenario 3: User Registration")
	fmt.Println("For the user's date of birth, use a data type like time.Time.")
	fmt.Println("For account verification status, use bool.")
}
