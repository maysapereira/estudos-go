package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	fmt.Println("Quando é sábado?")
	hoje := time.Now().Weekday()
	switch time.Saturday {
	case hoje + 0:
		fmt.Println("Hoje")
	case hoje + 1:
		fmt.Println("Amanhã.")
	case hoje + 2:
		fmt.Println("Daqui a dois dias")
	default:
		fmt.Println("Daqui a muitos dias")
	}
}
}