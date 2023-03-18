package main

import (
	"fmt"
	"net/http"
	"os"
)

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
	getData()
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

func getData() {
	data, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	bs := make([]byte, 99999)
	data.Body.Read(bs)
	fmt.Println(string(bs))
}
