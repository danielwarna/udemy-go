package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 10)
	hand.print()
	fmt.Println("==========")
	remainingDeck.print()

	fmt.Println(cards.toString())

	hand.saveToDisk("cards.file")

	deck2 := newDeckFromFile("cards.file")

	fmt.Println("======== Deck2 ============")
	deck2.print()

	fmt.Println("======== Suffled ============")
	deck2.shuffle()
	deck2.print()
}

func newCard() string {
	return "Five of Diamonds"
}
