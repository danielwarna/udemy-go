package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// A deck which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Knight", "Queen", "King"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() {

	for i := range d {
		newPos := rand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]

	}
}

func (d deck) toString() string {
	ds := []string(d)

	st := strings.Join(ds, ",")

	return st
}

func newDeckFromString(data string) deck {
	var deckOfCards deck
	c := strings.Split(data, ",")

	for _, card := range c {
		deckOfCards = append(deckOfCards, card)
	}

	return deckOfCards
}

func (d deck) saveToDisk(filename string) error {
	mBytes := []byte(d.toString())
	return os.WriteFile(filename, mBytes, 0666)
}

func newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Failed to load file " + filename)
		fmt.Println(err)
		os.Exit(1)
	}

	return newDeckFromString(string(data))

}
