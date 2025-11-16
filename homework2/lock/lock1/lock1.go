package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				count++
				fmt.Println(count)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
}
