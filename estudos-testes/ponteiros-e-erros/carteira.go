package main

type Carteira struct {
	saldo int
}

func (c *Carteira) Depositar(quantidade int) {
	c.saldo += quantidade
}

func (c *Carteira) Saldo() int {
	return c.saldo
}
