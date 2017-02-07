package main

import (
	"fmt"
	"math"
)

type GetMetry interface {
	area() float64
	perim() float64
}

type ToString interface {
	string() string
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface.
// 传参仕只能使用值传递，不能传递指针。

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (r Rect) string() string {
	return fmt.Sprintf("Rect: with=%f, height=%f", r.width, r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g GetMetry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := Rect{3, 4}
	c := Circle{5}

	measure(r)
	fmt.Println(r.string())

	measure(c)
}
