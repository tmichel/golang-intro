package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	R uint
}

func main() {
	c := Circle{Point{1, 1}, 3}
	fmt.Println("x", c.X, "y", c.Y, "r", c.R)
}
