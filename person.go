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

func (personPointer *person) update(newFirstname string) {
	(*personPointer).firstname = newFirstname
}
