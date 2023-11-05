package main

import (
	"fmt"
)

// Function with single parameters of different types
func printValues(i int, f float64, s string, b bool, customStruct MyStruct) {
	fmt.Println("Integer:", i)
	fmt.Println("Float:", f)
	fmt.Println("String:", s)
	fmt.Println("Boolean:", b)
	fmt.Println("Custom Struct:", customStruct)
}

// Function with multiple parameters of different types
func calculateSumAndProduct(a, b int, x, y float64, s1, s2 string) (int, float64) {
	sum := a + b
	product := x * y
	return sum, product
}

// Function with variadic parameters (integers)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function that accepts a slice as a variadic parameter
func processSlice(data ...string) {
	for _, item := range data {
		fmt.Println("Item:", item)
	}
}

// Struct to demonstrate custom data type
type MyStruct struct {
	Field1 int
	Field2 string
}

func main() {
	customData := MyStruct{Field1: 42, Field2: "Hello, Go!"}

	// Testing single parameter function
	fmt.Println("Single Parameter Function:")
	printValues(42, 3.1416, "Go is fun", true, customData)
	fmt.Println()

	// Testing multiple parameter function
	fmt.Println("Multiple Parameter Function:")
	sumResult, productResult := calculateSumAndProduct(10, 20, 2.5, 3.0, "Hello", "World")
	fmt.Println("Sum:", sumResult)
	fmt.Println("Product:", productResult)
	fmt.Println()

	// Testing variadic parameter function
	fmt.Println("Variadic Parameter Function:")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of Numbers:", sum(numbers...))
	fmt.Println()

	// Testing variadic parameter function with a slice
	fmt.Println("Variadic Parameter Function with Slice:")
	dataSlice := []string{"apple", "banana", "cherry"}
	processSlice(dataSlice...)
}
