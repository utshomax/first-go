package main

import (
	"fmt"
	"io"
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

//My custom logwriter

type logWritter struct{}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func getData() {
	data, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	// bs := make([]byte, 99999)
	// data.Body.Read(bs)
	// fmt.Println(string(bs))
	//addding writter
	lw := logWritter{}
	io.Copy(lw, data.Body)
}
