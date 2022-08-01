package main

import (
	"net/http"
	"time"
)

func Corredor(a, b string) (vencedor string) {
	duracaoA := medirTempoDeResposta(a)
	duracaoB := medirTempoDeResposta(b)

	if duracaoA < duracaoB {
		return a
	}

	return b
}

func medirTempoDeResposta(URL string) time.Duration {
	inicio := time.Now()
	http.Get(URL)
	return time.Since(inicio)
}
