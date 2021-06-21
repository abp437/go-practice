package langintro

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

// Notes:
// A Struct is just a collection of data types in Go.
// Struct is like a Class, it's a blueprint of an object in Go.
// It uses pointers extensively to pass on structs and modify and access it's properties and functions.
// We can add methods to our custom structs with Receiver Functions.
