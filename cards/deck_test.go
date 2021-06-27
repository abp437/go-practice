package cards

import "testing"

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	// Deck Size
	if len(d) != 52 {
		t.Errorf("Expected Deck length of 52, but got %v", len(d))
	}

	// First Card in Deck
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spade sto be the first card in the deck, instead got %v", d[0])
	}

	// Last Card in Deck
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs to be the last card in the deck, instead got %v", d[len(d)-1])
	}
}

// Notes:
// You can write tests around every function
// We need to export all the test functions to be accessible on terminal
