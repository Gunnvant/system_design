package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	counter := packItems(0)
	fmt.Printf("total items packed %d", counter)
}

func packItems(totalItems int32) int32 {
	const workers = 2
	const itemsPerWorker = 1000
	var wg sync.WaitGroup
	itemsPacked := int32(0)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workderId int) {
			defer wg.Done()
			for j := 0; j < itemsPerWorker; j++ {
				atomic.AddInt32(&itemsPacked, 1)
			}
			atomic.SwapInt32(&totalItems, itemsPacked)
		}(i)

	}
	wg.Wait()
	return totalItems
}
