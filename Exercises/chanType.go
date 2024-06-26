package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)

	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println("DADADADA")
			fmt.Println(i)
		} 
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		ch <- 244
		ch <- 27556
	
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}