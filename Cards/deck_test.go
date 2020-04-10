package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	const dLen = 52
	const first = "Ace of Spades"
	const last = "King of Clubs"
	d := newDeck()

	if len(d) != dLen {
		t.Errorf("Expected deck of %d cards but got %d", dLen, len(d))
	}

	if d[0] != first {
		t.Errorf("Expected the first card to be %v but got %v", first, d[0])
	}

	if d[dLen-1] != last {
		t.Errorf("Expeced the last card to be %v but got %v", last, d[dLen-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	const testFileName = "$$_decktesting"
	os.Remove(testFileName)
	d := newDeck()
	d.saveToFile(testFileName)

	ld := newDeckFromFile(testFileName)
	if len(ld) != len(d) {
		t.Errorf("Expected deck length to be %v but got %v", len(d), len(ld))
	}

	os.Remove(testFileName)
}
