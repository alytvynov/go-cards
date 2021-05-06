package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "1 ♠ - Spades" {
		t.Errorf("Expected `1 ♠ - Spades`, but goot %v", d[0])
	}

	if d[len(d)-1] != "5 ♣ - Clubs" {
		t.Errorf("Expected `5 ♣ - Clubs`, but goot %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	filename := "_decktesting.txt"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cars un deck, got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
