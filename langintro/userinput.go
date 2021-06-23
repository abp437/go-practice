package langintro

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Exported Functions:

// CreateBill creates a bill through user input and then sends it back to the user
func CreateBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := cliInput(reader, "Enter your name: ")
	b := newBill(name)
	promptOptions(reader, &b)
	return b
}

// Package Functions:

func cliInput(r *bufio.Reader, p string) (string, error) {
	fmt.Print(p)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(r *bufio.Reader, b *Bill) {
	opt, _ := cliInput(r, "Choose Option (a - add item, s - save bill, t - add tip) \n")
	switch strings.ToLower(opt) {
	case "a":
		fmt.Println("Adding item")
	case "s":
		fmt.Println("Saving bill")
	case "t":
		fmt.Println("Adding tip")
	default:
		fmt.Println("Enter a correct option, Jabroni")
	}
}

// Notes:
// In idomatic Go, we keep the exported functions at the top
// Then comes the most used or important private package scoped functions, then comes the normal private package scoped functions
// Here we use pointers extensively for passing as arguments since we don't want to create multiple readers, since it will waste memory
