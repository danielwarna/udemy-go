package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck lenght to be 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be ace of spades %v", len(d[0]))
	}
}

func TestShuffleAndDeal(t *testing.T) {

	d := newDeck()

	d1, d2 := deal(d, 10)

	if len(d1) != 10 {
		t.Errorf("Expected deal 10 to return deck of len 10, got %v", len(d[0]))
	}
	if len(d2) != len(d)-10 {
		t.Errorf("Expected deal 10 to return deck of len %v, got %v", len(d)-10, len(d[0]))
	}

	first := d1[0]
	second := d1[1]
	d1.shuffle()
	if first == d1[0] && second == d1[1] {

		t.Errorf("Expected shuffle to actually shuffle")
	}
}

func TestLoadDeck(t *testing.T) {

	d := newDeckFromFile("test_cards.file")
	if d[0] != "Ace of Heats" {
		t.Errorf("Expected the first card to be Ace of Heats %v", len(d[0]))
	}
}

func TestSaveAndLoadDeck(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()
	d.saveToDisk("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
