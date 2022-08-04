package main

import (
	"fmt"
	"time"
)

func numeros() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func letras() {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
	}
}

func main() {
	go numeros()
	go letras()
	time.Sleep(5 * time.Second)
	fmt.Println("Fim da execução!")
}
