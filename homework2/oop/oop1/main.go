package main

import (
	"fmt"
)

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	a int
	b int
}

func (r *Rectangle) Area() {
	fmt.Println(r.Height * r.Width)
}
func (r *Rectangle) Perimeter() {
	fmt.Println(r.Height * r.Width * r.Height * r.Width)
}

func (c *Circle) Area() {
	fmt.Println(c.a * c.b)
}
func (c *Circle) Perimeter() {
	fmt.Println(c.a * c.b)
}

func main() {
	r := &Rectangle{Width: 3, Height: 5}
	c := &Circle{a: 1, b: 2}
	r.Perimeter()
	r.Area()

	c.Perimeter()
	c.Area()
}
