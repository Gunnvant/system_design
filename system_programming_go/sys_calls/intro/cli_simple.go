package main

import (
	"fmt"
	"os"
)

func main() {
	words := os.Args[1:]
	if len(words) == 0 {
		fmt.Println("Please provide some words as arguments.")
		os.Exit(1)
	}
	for _, w := range words {
		if len(w)%2 == 0 {
			fmt.Fprintf(os.Stdout, "word %s is even\n", w)
		} else {
			fmt.Fprintf(os.Stdout, "word %s is odd\n", w)
		}
	}

}
