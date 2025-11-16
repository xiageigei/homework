package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go printOdd(10)
	go printEven(10)
	wg.Wait()
}

func printOdd(param int) {
	for i := 1; i <= param; i += 2 {
		fmt.Println("输出奇数:", i)
	}
	wg.Done()
}

func printEven(param int) {
	for i := 0; i <= param; i += 2 {
		fmt.Println("输出偶数:", i)
	}
	wg.Done()
}
