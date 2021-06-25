package cards

import "fmt"

// Deck is actually a deck of 52 cards, which is essentially a slice of strings
// Declaring a custom Struct with a single field appears like this.
type Deck []string

// Print prints all the cards in a Deck
func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
