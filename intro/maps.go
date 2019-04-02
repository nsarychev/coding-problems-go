package main

import (
	"fmt"
)

func maps() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	_, exists := lookup["vegeta"]

	// prints 0, false
	// 0 is the default value for an integer
	fmt.Println(exists)

	total := len(lookup)
	fmt.Println(total)
	delete(lookup, "goku")
	lookup = make(map[string]int, 100)
}
