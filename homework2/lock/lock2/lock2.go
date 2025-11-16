package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func taskAdd(val *int32) {
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(val, 1)
	}
}

func main() {
	var wg sync.WaitGroup
	c := int32(0)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			taskAdd(&c)
		}()
	}
	wg.Wait()
	fmt.Println(c)
}
