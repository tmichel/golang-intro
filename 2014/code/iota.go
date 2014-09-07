package main

import (
	"fmt"
	"time"
)

// START OMIT
const (
	zero = iota
	one  = iota
	two  = iota
)

const (
	pow0 time.Duration = 1 << iota
	pow1
	pow2
)

const (
	StatusA = iota
	StatusB
	_ // skip the value '2'
	StatusC
)

// END OMIT

func main() {
	fmt.Println("zero", zero, "one", one, "two", two)
	fmt.Printf("%s %s %s %s %s %s %T\n", "pow0", pow0, "pow1", pow1, "pow2", pow2, pow0)
	fmt.Println("StatusA", StatusA, "StatusB", StatusB, "StatusC", StatusC)
}
