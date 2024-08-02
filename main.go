package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	fmt.Println(cards)
}
