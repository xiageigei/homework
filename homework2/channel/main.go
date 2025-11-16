package main

import "fmt"

func main() {
	ch := make(chan int, 100)
	go sendOnly(ch)
	receive(ch)
}

func sendOnly(ch chan<- int) {
	for i := 0; i < 100; i++ {
		fmt.Println("发送", i)
		ch <- i
	}
	defer close(ch)
}

func receive(ch <-chan int) {
	for v := range ch {
		fmt.Println("接收", v)
	}
}
