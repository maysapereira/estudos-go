package main

import (
	"reflect"
	"testing"
)

func TestPercorre(t *testing.T) {

	casos := []struct {
		Nome              string
		Entrada           interface{}
		ChamadasEsperadas []string
	}{
		{
			"Struct com um campo string",
			struct {
				Nome string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct com dois campos tipo string",
			struct {
				Nome   string
				Cidade string
			}{"Maysa", "Rio de Janeiro"}, []string{"Maysa", "Rio de Janeiro"},
		},
		{
			"Struct sem campo tipo string",
			struct {
				Nome  string
				Idade int
			}{"Maysa", 22}, []string{"Maysa"},
		},
	}

	for _, teste := range casos {
		t.Run(teste.Nome, func(t *testing.T) {
			var resultado []string

			percorre(teste.Entrada, func(entrada string) {
				resultado = append(resultado, entrada)
			})

			if !reflect.DeepEqual(resultado, teste.ChamadasEsperadas) {
				t.Errorf("resultado %v, esperado %v", resultado, teste.ChamadasEsperadas)
			}
		})
	}
}
