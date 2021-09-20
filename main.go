package main

import "fmt"

type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Triangle struct {
	base, height float64
}

func (s Triangle) area() float64 {
	return (s.base * s.height) / 2
}

func printArea(s Shape) {
	fmt.Print(s.area())
	fmt.Print("\n")
}

func main() {
	printArea(Square{5})
	printArea(Triangle{5, 100})
}
