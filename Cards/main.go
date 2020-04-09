package main

func main() {
	cards := newDeck()

	// var hand deck
	// hand, cards = cards.deal(5)
	// hand.print()
	// fmt.Println(hand.toString())
	// fmt.Println("----------")
	cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())
}
