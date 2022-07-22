package main

import "fmt"

func main() {
	var x int = 100
	var y *int = &x

	fmt.Println(x, y)

	fmt.Println(x, *y)

}
