package main

import "fmt"

type squre struct {
	hand float64
}

// A perfect triange
type triange struct {
	hand float64
}

type shape interface {
	getArea() float64
}

func main() {
	squre := squre{
		hand: 4.50,
	}
	triange := triange{
		hand: 7.90,
	}
	printArea(squre)
	printArea(triange)
}

func (s squre) getArea() float64 {
	return s.hand * s.hand
}

func (t triange) getArea() float64 {
	return t.hand * 3
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
