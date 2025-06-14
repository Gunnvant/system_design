package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Mutex{}
	counter := packitems(&m, 0)
	fmt.Printf("total items packed %d", counter)
}

func packitems(m *sync.Mutex, totalItems int) int {
	const workers = 2
	const itemsPerWorker = 1000
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workderID int) {
			defer wg.Done()
			for j := 0; j <= itemsPerWorker; j++ {
				m.Lock()
				itemsPacked := totalItems
				itemsPacked++
				totalItems = itemsPacked
				m.Unlock()
			}
		}(i)
	}
	wg.Wait()
	return totalItems

}
