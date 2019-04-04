package main

import (
	"fmt"
	"os"
)

func deferFile() {
	file, err := os.Open("a_file_to_read")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // explicitly release
}
