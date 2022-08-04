package main

import (
	"fmt"
	"time"
)

func numeros(n chan<- int) {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("Escrito no channel %d\n", i)
		time.Sleep(time.Millisecond * 90)
	}

	fmt.Println("Fim da escrita")
	close(n)
}

func main() {
	cn := make(chan int, 3)
	go numeros(cn)

	for v := range cn {
		fmt.Printf("lido do channel: %d\n", v)
		time.Sleep(time.Millisecond * 180)
	}

	fmt.Println("Fim da execução!")
}
