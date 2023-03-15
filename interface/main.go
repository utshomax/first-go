package main

import "fmt"

type englistBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englistBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
func (eb englistBot) getGreeting() string {
	return "hello word !"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
