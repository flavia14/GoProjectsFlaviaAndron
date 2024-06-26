package main

import (
	"fmt"
	"time"
)
func main() {
	go sayHello()
	go sayHello2()
	time.Sleep(100 * time.Millisecond)

}


func sayHello() {
	fmt.Println("hello")
}

func sayHello2() {
	fmt.Println("hellodddd")
	time.Sleep(200 * time.Millisecond)

}