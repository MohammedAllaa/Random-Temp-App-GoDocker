package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random temperature between 0 and 40
	temperature := rand.Intn(41) // 0–40 inclusive

	// Print the temperature
	fmt.Printf("Temperature: %d°C\n", temperature)

	// Use switch to print message based on temperature
	switch {
	case temperature < 10:
		fmt.Println("It's cold")
	case temperature >= 10 && temperature <= 25:
		fmt.Println("It's moderate")
	default:
		fmt.Println("It's warm")
	}
}