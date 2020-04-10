package main

import "testing"

func TestNewDeck(t *testing.T) {
	const dLen = 52
	d := newDeck()

	if len(d) != dLen {
		t.Errorf("Expected deck of %d cards but got %d", dLen, len(d))
	}
}
