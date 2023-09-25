package main

import (
	"fmt"
)

func main() {
	word := "selamat malam"
	count := make(map[string]int);

	for _, char := range word {
		fmt.Printf("%c\n", char);
		count[string(char)]++;
	}

	// print the map of characters total occurences
	fmt.Printf("%v\n", count);
}