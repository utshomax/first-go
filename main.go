package main

func main() {
	cards := newDeck()
	hand, remianingCards := deal(cards, 5)
	hand.print()
	remianingCards.print()
}
