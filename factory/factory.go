package main

import (
	"estudos/estudos-go/estudos/factory/test"
	"fmt"
)

func main() {
	c := test.NewCar()

	c.DefineQtdRodas(-20)
	fmt.Println(c)
}
