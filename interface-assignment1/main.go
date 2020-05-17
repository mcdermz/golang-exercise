package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func main() {
	s := square{
		sideLength: 10.1,
	}
	t := triangle{
		height: 5.02,
		base:   7.33,
	}

	printArea(s)
	printArea(t)
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
