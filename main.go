package main

import "fmt"

func main() {
	cards := newDeck()
	// hand, remainginDeck := deal(cards, 5)

	// hand.print()
	// remainginDeck.print()

	fmt.Println(cards.toString())
}
