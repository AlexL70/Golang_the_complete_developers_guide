package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	tr := triangle{height: 2, base: 4}
	sq := square{sideLength: 3}
	fmt.Print(tr, " – this is a triangle its area is ")
	printArea(tr)
	fmt.Print(sq, " – this is a square its are is ")
	printArea(sq)
}
