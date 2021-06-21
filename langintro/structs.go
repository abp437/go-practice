package langintro

import "fmt"

// Bill is the object(struct) we want to use to create a customer's payment record
type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// NewBill is like a constructor, it creates a new bill wth initial values and returns the Bill
func NewBill(n string) Bill {
	b := Bill{
		name:  n,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Format receiver function/method takes in a Bill and then returns a formatted Bill
func (b *Bill) Format() string {
	fs := b.name + "'s Bill Breakdown:\n"
	total := 0.0
	// Cycle through items on Bill
	for k, v := range b.items {
		fs += fmt.Sprintf("%v: %v\n", k, v)
		total += v
	}
	// Add Tip
	if b.tip != 0 {
		fs += fmt.Sprintf("Tip: %0.2f\n", b.tip)
		total += b.tip
	}
	fs += fmt.Sprintf("Total: %0.2f\n", total)
	return fs
}

// UpdateTip receiver function/method adds an extra tip to a bill
func (b *Bill) UpdateTip(n float64) {
	b.tip = n
}

// AddItem receiver function/method adds an item to a bill
func (b *Bill) AddItem(name string, price float64) {
	b.items[name] = price
}

// Notes:
// A Struct is just a collection of data types in Go.
// Struct is like a Class, it's a blueprint of an object in Go.
// It uses pointers extensively to pass on structs and modify and access it's properties and functions.
// We can add methods to our custom structs with Receiver Functions.
// We take pointer receiver functions wherever we want to modify the original object itself.
// Pointers are automatically dereferenced in Receiver Functions.
