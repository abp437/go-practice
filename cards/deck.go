package cards

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

// Exported Structs:

// Deck is actually a deck of 52 cards, which is essentially a slice of strings
// Declaring a custom Struct with a single field appears like this.
type Deck []string

// Exported Functions:

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

// Deal takes out a hand of cards from a Deck and returns the hand and the leftover cards
func Deal(d Deck, size int) (Deck, Deck) {
	return d[:size], d[size:]
}

// Shuffle shuffles all the cards in a deck of cards
func (d Deck) Shuffle() {
	// Need this for unique generation of a random number every time
	// This generates a random number not only at compilation but also on every execution
	rand.Seed(time.Now().UnixNano())

	for i := range d {
		randNum := rand.Intn(len(d) - 1)
		d[i], d[randNum] = d[randNum], d[i]
	}
}

// SaveToFile takes in the contents of a Deck of cards and saves it to a file
func (d Deck) SaveToFile(filepath string) error {
	return ioutil.WriteFile(filepath, []byte(d.toString()), 0644)
}

// NewDeckFromFile reads the contents from a file and creates a Deck Struct and returns it
func NewDeckFromFile(filepath string) Deck {
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return []string(strings.Split(string(contents), "\n"))
}

// Package Functions:

func (d Deck) toString() string {
	return strings.Join(d, "\n")
}

// Notes:
// A Byte Slice is basically another way of representing a string of characters.
// Every element in that slice corresponds to ASCII value of the character.
