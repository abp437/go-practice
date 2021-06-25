package main

import (
	"go-practice/cards"
)

func main() {
	deck := cards.Deck{
		"Ace of Spades",
		"Five of Diamonds",
	}

	deck.Print()
}
