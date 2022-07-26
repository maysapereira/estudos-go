package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}
	t.Run("diz olá para uma pessoa", func(t *testing.T) {
		resultado := Ola("Maysa", "")
		esperado := "Olá, Maysa"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá, mundo quando uma string vazia é passada como parâmetro", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá em espanhol", func(t *testing.T) {
		resultado := Ola("Maysa", espanhol)
		esperado := "Hola, Maysa"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá em francês", func(t *testing.T) {
		resultado := Ola("Maysa", frances)
		esperado := "Bonjour, Maysa"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
