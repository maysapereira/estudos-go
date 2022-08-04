package main

import "fmt"

//Triangle

type triangle struct {
	size int
}

type coloredTriangle struct {
	triangle
	color string
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (t coloredTriangle) perimeter() int {
	return t.triangle.perimeter()
}

//Square
type square struct {
	size int
}

type coloredSquare struct {
	square
	color string
}

func (s square) perimeter() int {
	return s.size * 4
}

func (t coloredSquare) perimeter() int {
	return t.square.perimeter()
}

func main() {

	t := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Size:", t.size)

	s := coloredSquare{square{4}, "red"}
	fmt.Println("Perimeter (square):", s.perimeter())
	fmt.Println("Size:", s.size)
}
