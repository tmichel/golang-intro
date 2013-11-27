package main

import "fmt"

type Drawer interface {
	Draw()
}

type Point struct{ X, Y int }

func (Point) toS() string {
	return "Point"
}

func (Point) Draw() {
	fmt.Print(".")
}

type Circle struct {
	Point
	R int
}

func (Circle) toS() string {
	return "Circle"
}

type Cube int

func (Cube) Draw() {
	fmt.Print("\u2588")
}

func DrawEverything(ds ...Drawer) {
	for _, d := range ds {
		d.Draw()
	}
}

func main() {
	c := Circle{Point{X: 10, Y: 0}, 1}

	fmt.Printf("%#v\n", c)
	c.X = 1
	fmt.Printf("%#v\n", c)

	fmt.Println(c.toS())
	fmt.Println(c.Point.toS())

	cube := Cube(1)
	p := Point{}

	DrawEverything(p, cube)
	fmt.Println()

}
