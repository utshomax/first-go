package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("WebSite Status Checker")
	links := []string{
		"https://google.com",
		"https://fb.com",
		"https://utsabutsho.me",
		"https://stackoverflow.com",
		"https://github.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	for {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("DOWN --> ", link)
		c <- "May be Down"
		return
	}
	c <- "Yup ! It's Up."
	fmt.Println("UP -->", link)

}
