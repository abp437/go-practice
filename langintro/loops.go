package langintro

import "fmt"

// Exported Functions:

// NormalLoop prints the indexes until a provided count not including the count itself
func NormalLoop(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("Normal For Loop: %v \n", i)
	}
}

// RangeLoop prints a slice of strings through a loop
func RangeLoop() {
	names := []string{"Mario", "Luigi", "Bowser", "Peach"}

	for _, val := range names {
		fmt.Printf("Range For Loop: %v \n", val)
	}
}

// Notes:
// "break" keyword breaks out of the loop completely
// "continue" keyword skips the current iteration of the loop
