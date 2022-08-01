package main

import (
	"fmt"
	"time"
)

func fun(value string) {
	for i := 0; i < 3; i++ {
		fmt.Println(value)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	fun("Direct call")

	// go fun("goroutine - 1")

	// fgx := fun
	// go fgx("goroutine - 2")

	// go func() {
	// 	fun("goroutine - 3")
	// }()

	time.Sleep(5 * time.Millisecond)

	fmt.Println("Done...")
}
