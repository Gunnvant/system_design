package main

import (
	"fmt"
	"sync"
)

func throwBallR(color string, ball chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Throwing ball of %s color \n", color)
	ball <- color
}

func receiveBallR(color string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Received ball of %s color\n", color)
}

func main() {
	var wg_throwers sync.WaitGroup
	var wg_receivers sync.WaitGroup
	ball := make(chan string)
	colors := []string{"red", "blue", "green", "white", "black"}
	for _, color := range colors {
		wg_throwers.Add(1)
		go throwBallR(color, ball, &wg_throwers)
	}
	go func() {
		wg_throwers.Wait()
		close(ball)
	}()
	for col := range ball {
		wg_receivers.Add(1)
		go receiveBallR(col, &wg_receivers)
	}
	wg_receivers.Wait()
}
