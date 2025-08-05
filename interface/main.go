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
	Height float64
	Width  float64
}

func (r rect) area() float64 {
	return r.Height * r.Width
}

func (r rect) perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

type circle struct {
	Radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func print_shape(s shape) {
	fmt.Printf("Area of %v: %.2f\n", s, s.area())
	fmt.Printf("Perimeter of %v : %.2f\n", s, s.perimeter())
}

func test(s shape) {
	fmt.Println("===============================")
	print_shape(s)
	fmt.Println("===============================")
}

func main() {
	test(rect{
		Height: 4,
		Width:  3,
	})

	test(circle{
		Radius: 5,
	})

	test(rect{
		Width:  9000,
		Height: 0.0005,
	})

}
