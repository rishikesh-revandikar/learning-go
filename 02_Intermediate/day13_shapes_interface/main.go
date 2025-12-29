package main

import (
	"fmt"
	"math"
)

// Shape interface defines the required behavior
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

// Area implementation for Circle: Area = πr²
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Area implementation for Rectangle: Area = w * h
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}

	// Both types fit into the Shape interface
	shapes := []Shape{c, r}

	for _, shape := range shapes {
		printArea(shape)
	}
}