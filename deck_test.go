package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	expected := "2 of Diamonds"
	actual := d[0]
	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)
	expected := 52
	actual := len(loadedDeck)
	if actual != expected {
		t.Errorf("Expected deck length of %v, but got %v", expected, actual)
	}

	os.Remove(filename)
}
