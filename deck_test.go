package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}
	if d[0] != "Ace of Speades" {
		t.Errorf("Expected 1st card as Ace of Speads but got %v", d[0])
	}
	if d[len(d)-1] != "Ten of Clubs" {
		t.Errorf("Expected last card as Ten of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	d := newDeck()
	os.Remove("_decktesting")
	d.saveToFile("_decktesting")
	deck, _ := readDeckFromFile("_decktesting")
	if len(deck) != 40 {
		t.Errorf("Expected 40 card but got %v", len(deck))
	}
	os.Remove("_decktesting")
}
