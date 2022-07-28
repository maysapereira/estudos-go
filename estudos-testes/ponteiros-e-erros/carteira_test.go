package main

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	confirmaSaldo := func(t *testing.T, carteira Carteira, esperado Bitcoin) {
		t.Helper()
		resultado := carteira.Saldo()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Sacar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		carteira.Sacar(10)
		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Sacar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Sacar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)

		if erro == nil {
			t.Error("Esperava um erro mas nenhum ocorreu")
		}
	})
}
