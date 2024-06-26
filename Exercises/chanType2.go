package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println(time.Now())
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for {
			if i, ok := <- ch; ok {
				fmt.Println(i)
			} else {
				break
			}
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