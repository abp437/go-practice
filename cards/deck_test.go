package cards

import "testing"

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 52 {
		t.Errorf("Expect Deck length of 52, but got %v", len(d))
	}
}

// Notes:
// You can write tests around every function
// We need to export all the test functions to be accessible on terminal
