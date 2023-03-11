package main

import "fmt"

type deck []string

// func main() {
// 	fmt.Println("deck")
// }

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) addCard(newcard string) deck {
	d = append(d, newcard)
	fmt.Println("New card Added", newcard)
	return d
}

func newDeck() deck {
	cardSlices := []string{"Speades", "Diamonds", "Hearts", "Clubs"}
	cardvalues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	cards := deck{}

	for _, value := range cardvalues {
		for _, slice := range cardSlices {
			cards = append(cards, value+" of "+slice)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[:handSize]
}
