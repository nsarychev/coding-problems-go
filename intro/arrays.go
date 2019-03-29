package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func arraysAndSlices() {
	arr := [5]int{9001, 9333, 212, 33}
	arr[4] = 339
	slice := []int{1, 4, 293, 4, 9}
	arrayRange(arr)
	sliceRange(slice)
	madeSlice := make([]int, 0, 10)  // length 0, capacity 10 (underlying array size)
	madeSlice = append(madeSlice, 5) // add 5 as a first element, and returns (re-assign to the var)
	sliceRange(madeSlice)

	reSliced := madeSlice[0:8] // re-slice
	reSliced[7] = 9033
	sliceRange(reSliced)

	testSliceAutoGrowth()

	scores := make([]int, 5)      // makes with cap and length of 6
	scores = append(scores, 9332) // appends at the end as 6th element since it is already length of 5
	sliceRange(scores)

	var nilSlice []string
	nilSlice = append(nilSlice, "no more nil")
	stringSliceRange(nilSlice)

	haystack := "the spice must flow"
	strings.Index(haystack[5:], " ") // first space in a string after first 5 characters

	scores = []int{1, 2, 3, 4, 5}
	scores = scores[:len(scores)-1] // all except for last
	fmt.Println(scores)

	scores = removeAtIndex(scores, 2)
	fmt.Println(scores)

	// copy
	scores = make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst, scores[:5])
	fmt.Println(worst)\
}

func arrayRange(a [5]int) {
	for index, value := range a {
		fmt.Printf("%d : %d\n", index, value)
	}
}

func sliceRange(a []int) {
	fmt.Printf("Capacity: %d\n", cap(a))
	for index, value := range a {
		fmt.Printf("%d : %d\n", index, value)
	}
}

func stringSliceRange(a []string) {
	fmt.Printf("Capacity: %d\n", cap(a))
	for index, value := range a {
		fmt.Printf("%d : %s\n", index, value)
	}
}

func testSliceAutoGrowth() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Printf("Initial capacity: %d\n", c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		// if our capacity has changed,
		// Go had to grow our array to accommodate the new data
		if cap(scores) != c {
			c = cap(scores)
			fmt.Printf("New capacity: %d\n", c)
		}
	}
}

// won't preserve order
func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	//swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
