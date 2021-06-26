package cards

import "fmt"

// Deck is actually a deck of 52 cards, which is essentially a slice of strings
// Declaring a custom Struct with a single field appears like this.
type Deck []string

// NewDeck creates a new deck of cards and returns it
func NewDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, s := range cardSuits {
		for _, v := range cardValues {
			cards = append(cards, v+" of "+s)
		}
	}
	return cards
}

// Print prints all the cards in a Deck
func (d Deck) Print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
