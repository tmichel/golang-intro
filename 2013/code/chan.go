package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func(ch chan<- int) {
		fmt.Println("Go routine called.")
		time.Sleep(1 * time.Second)
		c <- 42
	}(c)

	fmt.Println(<-c)
}
