package langintro

import "math"

// HigherOrderFunction accepts a function as an argument
func HigherOrderFunction(name string, higherFunction func(string)) {
	higherFunction(name)
}

// FunctionWithReturn just calculates the area of a circle and returns it
func FunctionWithReturn(r float64) float64 {
	return math.Pi * r * r
}

// Notes:
// Higher Order functions accept named as well as anonymous functions as arguments
// Such functions are used very much in implementing Closures(Private Variables)
