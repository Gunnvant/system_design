package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go func() {
		c <- "message" // sending to channel
	}()
	fmt.Println(<-c) //receiving channel
}
