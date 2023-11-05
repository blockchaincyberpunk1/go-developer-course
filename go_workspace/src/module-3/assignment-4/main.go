// main.go
package main

import (
	"assignment-4/mathutil"
	"fmt"
)

func main() {
	a, b := 10, 5

	// Use functions from the mathutil package
	sum := mathutil.Add(a, b)
	diff := mathutil.Subtract(a, b)
	prod := mathutil.Multiply(a, b)
	quotient, err := mathutil.Divide(a, b)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Sum: %d\n", sum)
		fmt.Printf("Difference: %d\n", diff)
		fmt.Printf("Product: %d\n", prod)
		fmt.Printf("Quotient: %d\n", quotient)
	}
}
