package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	t.Run("retorna um erro se o servidor não responder dentro de 10s", func(t *testing.T) {
		servidorA := criarServidorComAtraso(11 * time.Second)
		servidorB := criarServidorComAtraso(12 * time.Second)

		defer servidorA.Close()
		defer servidorB.Close()

		_, err := Corredor(servidorA.URL, servidorB.URL)

		if err == nil {
			t.Error("esperava um erro, mas não obtive um")
		}
	})
}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))
}
