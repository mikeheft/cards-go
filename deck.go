package main

import (
	"fmt"
	"strconv"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	cardValues := makeCardSlices()

	for _, suit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, fmt.Sprintf("%v of %s", cardValue, suit))
		}
	}

	return cards
}

// d is similar to 'this' or 'self' in other languages
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func makeCardSlices() []string {
	faceCards := []string{"Ace", "King", "Queen", "Jack"}
	numberCards := []string{}

	for i := 2; i <= 10; i++ {
		numberCards = append(numberCards, strconv.Itoa(i))
	}
	return append(numberCards, faceCards...)
}
