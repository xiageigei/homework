package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tasks := []Task{
		func() { time.Sleep(2 * time.Second) },
		func() { time.Sleep(3 * time.Second) },
		func() { time.Sleep(4 * time.Second) },
	}
	Schedul(tasks)
}

type Task func()

func Schedul(tasks []Task) {
	wg := sync.WaitGroup{}

	for i, task := range tasks {
		wg.Add(1)
		go func(task Task, index int) {
			defer wg.Done()
			start := time.Now()
			task()
			end := time.Since(start)
			fmt.Println(index, end)
		}(task, i)
	}
	wg.Wait()
}
