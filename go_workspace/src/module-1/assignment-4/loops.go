package main

import "fmt"

func main() {
	// List of names
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}

	// Iterate through the list and print each name
	fmt.Println("List of Names:")
	for _, name := range names {
		fmt.Println(name)
	}
}
