package main

import (
	"bytes"
	"testing"
)

func TestContagem(t *testing.T) {
	buffer := &bytes.Buffer{}

	Contagem(buffer)

	resultado := buffer.String()
	esperado := `321Go!`

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
