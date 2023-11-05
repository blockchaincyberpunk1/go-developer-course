package main

import "fmt"

func main() {
	// Scenario a: Temperature Conversion
	celsiusTemp := 20.0 // Use float64 for precision
	fahrenheitTemp := (celsiusTemp * 9 / 5) + 32
	fmt.Printf("Task a: %.2f°C is equal to %.2f°F\n", celsiusTemp, fahrenheitTemp)

	// Scenario b: User Information
	name := "John Doe"          // Use string for name
	age := 30                   // Use int for age
	email := "john@example.com" // Use string for email address
	fmt.Printf("Task b: Name: %s\nAge: %d\nEmail: %s\n", name, age, email)

	// Scenario c: Inventory Management
	productName := "Laptop" // Use string for product name
	price := 999.99         // Use float64 for price
	quantityInStock := 50   // Use int for quantity
	fmt.Printf("Task c: Product: %s\nPrice: $%.2f\nQuantity in Stock: %d\n", productName, price, quantityInStock)

	// Scenario d: Financial Calculation (Compound Interest)
	principal := 1000.0 // Use float64 for principal amount
	rate := 0.05        // Use float64 for interest rate
	timeInYears := 5    // Use int for time in years
	finalAmount := principal * (1 + rate) * float64(timeInYears)
	fmt.Printf("Task d: Principal: $%.2f\nRate: %.2f\nTime (years): %d\nFinal Amount: $%.2f\n", principal, rate, timeInYears, finalAmount)

	// Scenario e: Geometry (Rectangle)
	length := 10.0 // Use float64 for length
	width := 5.0   // Use float64 for width
	area := length * width
	perimeter := 2 * (length + width)
	fmt.Printf("Task e: Length: %.2f\nWidth: %.2f\nArea: %.2f\nPerimeter: %.2f\n", length, width, area, perimeter)
}
