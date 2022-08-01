package main

import (
	"fmt"
	"net/http"
	"time"
)

func Corredor(a, b string) (vencedor string, erro error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("tempo limite de espera excedido para %s e %s", a, b)
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		ch <- true
	}()

	return ch
}
