package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func errorsRun() {
	arr := []string{"blah", "7"}
	if len(arr) != 2 {
		os.Exit(1) // dead code, just an example
	}

	n, err := strconv.Atoi(arr[1])
	if err != nil {
		fmt.Println("not a valid number")
	} else {
		fmt.Println(n)
	}

	analyze(1)
	analyze(0)
}

func analyze(i int) {
	error := isGteOne(i)
	if error == nil {
		fmt.Println("Yes")
	} else {
		fmt.Println(error)
	}
}

// Ypu can define a custom error interface
/* type error interface {
	Error() string
} */

func isGteOne(count int) error {
	if count < 1 {
		return errors.New("No")
	}
	return nil
}
