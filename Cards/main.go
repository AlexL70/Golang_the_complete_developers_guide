package main

import "fmt"

func main() {
	cards := newDeck()

	var hand deck
	hand, cards = cards.deal(5)
	hand.print()
	fmt.Println("----------")
	cards.print()
}
