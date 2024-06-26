package main

import (
	"fmt"
	"sync"
)

var (
	counter int32
	wg sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wg.Add(3)

	go increment("Flavia")
	go increment("Flavia Andron")
	go increment("Calina")

	wg.Wait()
	fmt.Println("Counter: ", counter)
}

func increment (lang string) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			counter++
		}
		mutex.Unlock()
	}
}