package main

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}

		carteira.Depositar(Bitcoin(10))

		resultado := carteira.Saldo()

		esperado := Bitcoin(10)

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	})

	t.Run("Sacar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}

		carteira.Sacar(Bitcoin(10))

		resultado := carteira.Saldo()

		esperado := Bitcoin(10)

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	})

}
