/*
	Cards

	newDeck 		: Create a list of playing cards. Essentially an array of strings
	print			: Log out the contents of a deck of cards
	shuffle			: Shuffles all the cards in a deck
	deal			: Create a 'hand' of cards
	saveToFile		: Save a list of cards to a file on the local machine
	newDeckFromFile	: Load a list of cards from the local machine

*/

package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}