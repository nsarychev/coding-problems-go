package main

import "fmt"

func collections() {
	var scores [10]int
	scores[0] = 339
	otherScores := []int{9001, 9333, 212, 33}
	println(otherScores[0])
	rangeA(otherScores)
}

func rangeA(a []int) {
	for index, value := range a {
		fmt.Printf("%d : %d\n", index, value)
	}
}
