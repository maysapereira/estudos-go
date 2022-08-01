package main

import (
	"net/http"
	"time"
)

func Corredor(a, b string) (vencedor string) {
	inicioA := time.Now()
	http.Get(a)
	duracaoA := time.Since(inicioA)

	inicioB := time.Now()
	http.Get(b)
	duracaoB := time.Since(inicioB)

	if duracaoA < duracaoB {
		return a
	}

	return b
}
