package main

import "fmt"

func main() {
	idades := make(map[string]uint8)
	idades["Maysa"] = 22
	idades["Kasper"] = 21
	idades["Iná"] = 24

	fmt.Println(idades)
}
