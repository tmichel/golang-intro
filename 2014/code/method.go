package main

import (
	"fmt"
)

type number int

func (n *number) inc() { // HL
	*n++
}

func main() {
	num := number(41)
	num.inc()
	fmt.Println(num)
}
