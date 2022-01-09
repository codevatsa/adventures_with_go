package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	areaOfTriangle := triangle{
		height: 30.5,
		base:   10.2,
	}
	areaOfSquare := square{
		sideLength: 6.3,
	}
	printArea(areaOfTriangle)
	printArea(areaOfSquare)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
