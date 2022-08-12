package test

import "fmt"

type carro struct {
	qtdRodas int
	Preco    float64
}

func (c *carro) DefinePreco(preco float64) {
	if preco < 5000 || preco > 1000000000 {
		fmt.Println("Não é permitido esse valor")
		return
	}
}

func (c *carro) DefineQtdRodas(qtdRodas int) {
	if qtdRodas < 0 || c.qtdRodas > 4 {
		fmt.Println("Não é permitido esse valor")
		return
	}
}

func NewCar() carro {
	return carro{4, 200000}
}
