package main

import (
	"fmt"
)

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	dog := Dog{
		Name:  "Karabaş",
		Breed: "Çoban Köpeği",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Hanzo",
		Color:         "Black",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says:", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Hav Hav"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Hugh"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}
