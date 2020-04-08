package main

func main() {
	cards := deck{"Ace of Diamonds", NewCard()}
	cards = append(cards, "Queene of Hearts")

	cards.print()
}

func NewCard() string {
	return "Five of Spades"
}
