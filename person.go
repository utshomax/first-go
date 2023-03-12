package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func newPerson(firstname string, lastname string, contact contactInfo) person {
	return person{firstname: firstname, lastname: lastname, contact: contact}
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) update(newFirstname string) {
	p.firstname = newFirstname
}
