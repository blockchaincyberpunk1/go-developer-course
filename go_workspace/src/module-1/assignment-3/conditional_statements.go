package main

import (
	"fmt"
)

func main() {
	// Input student's exam score (you can modify this value)
	score := 75

	// Check if the score meets the passing criteria
	if score >= 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
