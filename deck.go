package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	sliceOfStrings := []string(d)
	oneString := strings.Join(sliceOfStrings, ",")
	return oneString
}

func readDeckFromFile(filename string) (deck, error) {
	biteSlice, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	convertedDeck := strings.Split(string(biteSlice), ",")
	return deck(convertedDeck), error
}

func (d deck) shuffle() {

	//Truely random
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
