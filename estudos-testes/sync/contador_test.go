package main

import "testing"

func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes resultado no valor 3", func(t *testing.T) {
		contador := Contador{}
		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		if contador.Valor() != 3 {
			t.Errorf("resultado %d, esperado %d", contador.Valor(), 3)
		}
	})
}
