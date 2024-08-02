package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// This is to show the differences. IMO this should be a receiver
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("[ERROR] ", err)
		os.Exit(1)
	}

	str := string(bs)
	strArr := strings.Split(str, ",")

	return deck(strArr)
}

// mutates the caller
func (d deck) shuffle() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

	for i := range d {
		randIdx := r.Intn(len(d) - 1)
		d[i], d[randIdx] = d[randIdx], d[i]
	}
}

func (d deck) saveToFile(filename string) error {
	// 0666 == anyone can read/write
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func makeCardSlices() []string {
	faceCards := []string{"Ace", "King", "Queen", "Jack"}
	numberCards := []string{}

	for i := 2; i <= 10; i++ {
		numberCards = append(numberCards, strconv.Itoa(i))
	}

	return append(numberCards, faceCards...)
}
