package langintro

import "fmt"

// PrintHelloWorld prints "Hello, World!" to the Command Line
func PrintHelloWorld() {
	fmt.Println("Hello, World!")
}

// PrintHelloNinjas prints "Hello, Ninjas!" to the Command Line
func PrintHelloNinjas() {
	testVar := "Ninjas!"
	// "Sprintf" actually saves a formatted string to a variable, you can use it later
	formattedString := fmt.Sprintf("Hello, %v", testVar)
	fmt.Println("The saved formatted string is:", formattedString)
}

// PrintAgeName prints the age and name passed in to the function to the Command Line
func PrintAgeName(age int, name string) {
	// The format specifier for any variable to be print in default format is "%v"
	fmt.Printf("My age is %v, my name is %v\n", age, name)
}
