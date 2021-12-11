package main

import "testing"

func TestNewDeck(t *testing.T) { // 't' = test handler
	d := newDeck()

	if len(d) != 1600 {
		// Erorof, Logf
		t.Errorf("Expected deck length of 20, but got %v", len(d)) // Error Message
	}
}
