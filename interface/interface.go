package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height int32
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return (float64)(r.width * r.height)
}

func (r rect) perimeter() float64 {
	return (float64)(2*r.width + 2*r.height)
}

func (c circle) area() float64 {
	return (math.Pi * c.radius * c.radius)
}

func (c circle) perimeter() float64 {
	return (2 * math.Pi * c.radius)
}

func measure (s shape) {
	fmt.Printf("%T %+v\n",s, s)
	fmt.Printf("%v\n", s.area())
	fmt.Printf("%g\n\n", s.perimeter())
}

func main() {
	r := rect{width: 3,  height: 4}
	c := circle{radius:10}
	
	measure(r)
	measure(c)
}