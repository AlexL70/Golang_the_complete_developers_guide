package main

import "fmt"

func main() {
	card := NewCard()
	fmt.Println(card)
}

func NewCard() string {
	return "Five of Spades"
}
