package main

import "fmt"
import "time"

func main() {
  for {
    fmt.Println("I love logs")
    time.Sleep(1 * time.Second)
  }
}
