package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//	Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d[:], ",")
}

func fromString(s string) deck {
	return deck(strings.Split(s, ","))
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
