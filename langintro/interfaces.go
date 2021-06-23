package langintro

import (
	"fmt"
	"math"
)

// Exports:

// Square Struct
type Square struct {
	Length float64
}

// Circle Struct
type Circle struct {
	Radius float64
}

// Shape Interface
type Shape interface {
	area() float64
	circumference() float64
}

// ShapeInfo takes in a shape and prints the info of it
func ShapeInfo(s Shape) {
	fmt.Printf("Shape is of %T type, Area is %0.2f, Circumference is %0.2f \n", s, s.area(), s.circumference())
}

// Package Scope:

// Circle Methods
func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) circumference() float64 {
	return 2 * math.Pi * c.Radius
}

// Square Methods
func (s Square) area() float64 {
	return s.Length * s.Length
}

func (s Square) circumference() float64 {
	return s.Length * 4
}

// Notes:
// Interface groups types together based on their methods
// If you want a particular method to work across different Structs, then you have to use interfaces
// It's basically a common method designed to work across different Structs
// Here we can see that Circle and Square structs both implement "area" and "circumference" methods, so it actually implements the "shape" interface
// So, now these structs can be passed to all functions which accept a shape as an argument
// Interfaces don't have receiver functions, you can pass structs which implements the interface as an argument to functions which obey an interface
// All methods specified in an interface must be implemented to obey the interface
