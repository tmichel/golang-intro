package main

import (
	"fmt"
)

var (
	m    = mod
	rand = func() int {
		return 4 // chosen by fair dice roll
	}
)

func mod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	x, y := mod(10, 3)
	fmt.Println(x, y)
}
