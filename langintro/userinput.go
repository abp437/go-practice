package langintro

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Exported Functions:

// CreateBill creates a bill through user input and then sends it back to the user
func CreateBill() {
	reader := bufio.NewReader(os.Stdin)
	name, err := cliInput(reader, "Enter your name: ")
	if err != nil {
		log.Fatal(err)
	}

	b := newBill(name)
	promptOptions(reader, &b)
}

// Package Functions:

func cliInput(r *bufio.Reader, p string) (string, error) {
	fmt.Print(p)
	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(input), err
}

func promptOptions(r *bufio.Reader, b *Bill) {
	opt, err := cliInput(r, "Choose Option (a - add item, s - save the bill to a file, t - add tip) \n")
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(opt) {
	case "a":
		name, err := cliInput(r, "Item name \n")
		if err != nil {
			log.Fatal(err)
		}

		price, err := cliInput(r, "Item price \n")
		if err != nil {
			log.Fatal(err)
		}

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(r, b)
		}
		b.AddItem(name, p)
		fmt.Println("Item added - ", name, p)

		promptOptions(r, b)
	case "t":
		tip, err := cliInput(r, "Tip Amount ($) \n")
		if err != nil {
			log.Fatal(err)
		}

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(r, b)
		}
		b.UpdateTip(t)
		fmt.Println("Updated Tip - ", t)

		promptOptions(r, b)
	case "s":
		fmt.Println("Saving bill")
		b.saveBill()
	default:
		fmt.Println("Enter a correct option, Jabroni...")
		promptOptions(r, b)
	}
}

func (b *Bill) saveBill() {
	writeToFile(b.Format(), b.name)
}

// Notes:
// In idomatic Go, we keep the exported functions at the top
// Then comes the most used or important private package scoped functions, then comes the normal private package scoped functions
// Here we use pointers extensively for passing as arguments since we don't want to create multiple readers, since it will waste memory
