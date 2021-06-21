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
// Go's functions are Pass by Value, it makes copies of arguments when passed to functions
// But depending on the type of the argument passed, it either makes a copy of the actual data or a copy of a pointer to the data
// But in reality there's nothing called as Pass by Reference in Go
// It either copies actual data or a pointer
/*
	Pass by Value(Data copy) Types:
	Strings, ints, bools, floats, arrays and structs

	Pass by Reference(Pointer copy) Types:
	Slices, Maps and Functions

	Whenever we pass a Slice/Map/Function to a function as an argument we are actually passing a pointer.
	However it doesn't look like a pointer since it doesn't have a "*" in the function signature, but it
	implicitly is passing of a pointer.

	Q: What is the difference between the length and the capacity of a slice?
	A: A slice has both a length and a capacity.
	The length of a slice is the number of elements it contains.
	The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	The length and capacity of a slice "s" can be obtained using the expressions "len(s)" and "cap(s)".
*/
