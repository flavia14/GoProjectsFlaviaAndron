package main

import (
  "fmt"
  "time"
) 

func count(thing string){
  for i := 1; true; i++ {
    fmt.Println(i, thing)

    fmt.Scanln()
    // time.Sleep(time.Millisecond * 500)
  }

}

func main() {
 go count("da")
 go count("nu")

 time.Sleep(time.Second * 500)
}