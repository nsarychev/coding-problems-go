package main

import (
	"fmt"
)

func main() {
	power := getPower()
	goku := Saiyan{Name: "Goku"}
	goku.Power = power
	log(goku)
}

func getPower() int {
	return add(9001, 1)
}

func add(a int, b int) int {
	return a + b
}

func log(goku Saiyan) {
	fmt.Printf("Name: %s, Power %d\n", goku.Name, goku.Power)
}

type Saiyan struct {
	Name  string
	Power int
}
