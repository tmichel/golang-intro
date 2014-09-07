package main

import (
	"fmt"
)

type myString string

func greet(name myString) myString {
	return "Hi, " + name
}

func main() {
	var me string = "Tomi"
	fmt.Println(greet(me))
}
