package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 17:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}