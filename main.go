package main

func main() {
	cards := newDeck()
	hand, remainginDeck := deal(cards, 5)

	hand.print()
	remainginDeck.print()
}
