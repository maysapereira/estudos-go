package main

import "fmt"

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (s square) perimeter() int {
	return s.size * 4
}

func main() {
	t := triangle{3}
	s := square{4}

	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())
}
