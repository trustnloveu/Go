package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string

//! Receiver
func (d deck) print() {
	for i, card := range d { // 'd' can be replaced with 'this' or 'self'
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}
