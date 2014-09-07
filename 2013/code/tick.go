package main

import (
	"fmt"
	"time"
)

func main() {
	for t := range time.Tick(1 * time.Second) {
		fmt.Println(t.Second())
	}
}
