package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.print()
	// cards.shuffle()
	// cards.print()
	//cards.saveToFile("My cards")

	//Working with structs
	alex := newPerson("alex", "Doe", contactInfo{email: "utsho@test.com", zip: 00000})

	fmt.Println(alex)
}
