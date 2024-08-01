package main

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Ace of Diamonds")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
