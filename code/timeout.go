package main

import (
	"fmt"
	"time"
)

func main() {
	timeout(1500)
	timeout(2500)
}

func timeout(t int32) {
	done := make(chan int)

	go func(ch chan<- int) {

		// some long-running stuff
		time.Sleep(2 * time.Second)
		ch <- 1

	}(done)

	select {
	case d := <-done:
		fmt.Println("Done before timeout.", d)
	case <-time.After(time.Duration(t) * time.Millisecond):
		fmt.Println("Timed out")
	}
}
