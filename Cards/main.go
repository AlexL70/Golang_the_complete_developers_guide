package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", NewCard()}
	cards = append(cards, "Queene of Hearts")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func NewCard() string {
	return "Five of Spades"
}
