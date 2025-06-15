package main

import (
	"fmt"
	"sync"
)

func throwBall(color string, ball chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Throwing ball of %s color \n", color)
	ball <- color
}

func receiveBall(ball <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Received ball of %s color\n", <-ball)
}

func main() {
	var wg sync.WaitGroup
	ball := make(chan string)
	wg.Add(2)
	color := "blue"
	go throwBall(color, ball, &wg)
	go receiveBall(ball, &wg)
	wg.Wait()
}
