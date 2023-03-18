package main

import (
	"fmt"
	"net/http"
	"time"
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
	for l := range c {
		go func(newlink string) {
			time.Sleep(5 * time.Second)
			checkLink(newlink, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("DOWN --> ", link)
		c <- link
		return
	}
	c <- link
	fmt.Println("UP -->", link)
}
