package main

import (
	"fmt"
	"go-practice/cards"
)

func main() {
	deck := cards.NewDeck()
	hand, lftovr := cards.Deal(deck, 5)
	hand.Print()
	lftovr.Shuffle()
	lftovr.SaveToFile("./tmp/Deck.txt")
	fmt.Println(cards.NewDeckFromFile("./tmp/Deck.txt"))
}
