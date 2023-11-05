package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator with the current Unix timestamp to ensure different results every time.
	rand.Seed(time.Now().UnixNano())

	// Initialize the player's hit points (HP) to 100.
	playerHP := 100

	// Define an array of creature types the player may encounter.
	creatureTypes := []string{"goblin", "wolf", "troll"}

	// Continue the game as long as the player's HP is greater than 0.
	for playerHP > 0 {
		// Randomly select a creature from the creatureTypes array.
		creature := creatureTypes[rand.Intn(len(creatureTypes))]

		// Get the damage associated with the selected creature.
		creatureDamage := getCreatureDamage(creature)

		// Display a message to the player indicating the encounter.
		fmt.Printf("You encounter a %s!\n", creature)
		fmt.Printf("Choose an action: [attack/run/defend]: ")

		var choice string
		// Read the player's choice of action from the console.
		fmt.Scanln(&choice)

		// Process the player's chosen action.
		switch choice {
		case "attack":
			// Subtract the creature's damage from the player's HP and display the result.
			playerHP -= creatureDamage
			fmt.Printf("You attack the %s and lose %d HP.\n", creature, creatureDamage)
		case "run":
			// Determine if the player successfully runs away based on a 50% chance.
			if rand.Float64() < 0.5 {
				fmt.Printf("You successfully run away from the %s!\n", creature)
			} else {
				// If the run attempt fails, subtract the creature's damage from the player's HP.
				playerHP -= creatureDamage
				fmt.Printf("You tried to run but got caught by the %s and lose %d HP.\n", creature, creatureDamage)
			}
		case "defend":
			// Reduce the damage taken by half when defending.
			playerHP -= creatureDamage / 2
			fmt.Printf("You defend against the %s's attack and lose %d HP.\n", creature, creatureDamage/2)
		default:
			// Handle an invalid choice by informing the player.
			fmt.Println("Invalid choice. Please choose 'attack', 'run', or 'defend'.")
		}

		// Display the player's current HP after the action.
		fmt.Printf("Your HP: %d\n", playerHP)
	}

	// The game is over when the player's HP reaches 0 or below.
	fmt.Println("Game over! Your HP reached 0.")
}

// Function to retrieve the damage associated with a specific creature type.
func getCreatureDamage(creature string) int {
	switch creature {
	case "goblin":
		return 10
	case "wolf":
		return 15
	case "troll":
		return 20
	default:
		// Return 0 damage for unknown creatures.
		return 0
	}
}
