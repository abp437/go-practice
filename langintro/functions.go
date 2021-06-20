package langintro

import (
	"math"
	"strings"
)

// HigherOrderFunction accepts a function as an argument
func HigherOrderFunction(name string, higherFunction func(string)) {
	higherFunction(name)
}

// CircleArea just calculates the area of a circle and returns it
func CircleArea(r float64) float64 {
	return math.Pi * r * r
}

// PersonInitials takes in a name and returns the initials from the name
// This is a function with Multiple Return Values
func PersonInitials(n string) (string, string) {
	split := strings.Split(strings.ToUpper(n), " ")
	return split[0][:1], split[1][:1]
}

// Notes:
// Higher Order functions accept named as well as anonymous functions as arguments
// Such functions are used very much in implementing Closures(Private Variables)
