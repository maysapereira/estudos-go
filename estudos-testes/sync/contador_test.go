package main

import "testing"

func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes resultado no valor 3", func(t *testing.T) {
		contador := Contador{}
		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		verificaContador(t, contador, 3)

	})
}

func verificaContador(t *testing.T, resultado Contador, esperado int) {
	t.Helper()
	if resultado.Valor() != esperado {
		t.Errorf("resultado %d, esperado %d", resultado.Valor(), 3)
	}
}
