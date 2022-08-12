package test

import "fmt"

type carro struct {
	qtdRodas int
	Preco    float64
}

type moto struct {
	qtdRodas int
	Preco    float64
}

type Veiculo interface {
	PrintarQuantidadeRodas()
}

func (c *carro) PrintarQuantidadeRodas() {
	fmt.Println(c.qtdRodas)
}

func (m *moto) PrintarQuantidadeRodas() {
	fmt.Println(m.qtdRodas)
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

func NewVeiculo(qtdRodas int) Veiculo {
	if qtdRodas < 0 || qtdRodas > 4 {
		return nil
	}

	if qtdRodas == 2 {
		return &moto{qtdRodas: qtdRodas}

	} else if qtdRodas == 4 {
		return &carro{qtdRodas: qtdRodas}
	}

	return nil
}
