package main

import "fmt"

type triangle struct {
	size int
}

type square struct {
	size int
}

type coloredTriangle struct {
	triangle
	color string
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (s square) perimeter() int {
	return s.size * 4
}

func (t coloredTriangle) perimeter() int {
	return t.triangle.perimeter()
}

func main() {
	t := coloredTriangle{triangle{3}, "blue"}
	s := square{4}

	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Size:", t.size)

	fmt.Println("Perimeter (square):", s.perimeter())
}
