package main

import (
	"fmt"
)

func basics() {
	power := getPower()
	goku := &Saiyan{Person: &Person{"Gohan"}}
	goku.Power = power
	goku.Father = &Saiyan{
		Person: &Person{"Goku"},
		Father: nil,
	}
	goku.Super()
	goku.Father.Super()
	log(goku)
}

func getPower() int {
	return add(9001, 1)
}

func add(a int, b int) int {
	return a + b
}

func log(goku *Saiyan) {
	fmt.Printf("Name: %s, Power %d\n", goku.Name, goku.Power) // or can use goku.Person.Name
	goku.Introduce()
	goku.Person.Introduce()
	if goku.Father != nil {
		println("FATHER: ")
		log(goku.Father)
	}
}

type Saiyan struct {
	*Person
	Power  int
	Father *Saiyan
}

type Person struct {
	Name string
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

func (s *Saiyan) Introduce() {
	fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}
