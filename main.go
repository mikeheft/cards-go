package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
