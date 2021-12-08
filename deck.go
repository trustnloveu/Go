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
