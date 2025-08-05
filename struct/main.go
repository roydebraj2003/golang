package main

import "fmt"

type rect struct {
	Height int
	Width  int
}

//To define an function for the struct -

func (r rect) area() int {
	return r.Height * r.Width
}

func main() {
	r := rect{
		Height: 10,
		Width:  15,
	}
	fmt.Printf("The area of rect %v : %d\n", r, r.area())

}
