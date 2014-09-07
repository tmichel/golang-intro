package main

import "fmt"

func main() {
	s := "Hello, Go Budapest!"
	fmt.Printf("%s -> %s\n", s, reverse(s))
}

func reverse(s string) string {
	reversed := make([]rune, len(s))
	for i, r := range s {
		reversed[len(s)-i-1] = r
	}

	return string(reversed)
}
