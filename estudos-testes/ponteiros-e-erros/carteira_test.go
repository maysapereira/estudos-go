package main

import (
	"fmt"
	"testing"
)

func TestCarteira(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(10)

	resultado := carteira.Saldo()
	fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)
	esperado := 10

	if resultado != esperado {
		t.Errorf("resultado %d, esperado %d", resultado, esperado)
	}
}
