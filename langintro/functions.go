package langintro

// HigherOrderFunction accepts a function as an argument
func HigherOrderFunction(name string, higherFunction func(string)) {
	higherFunction(name)
}
