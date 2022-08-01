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
		{
			"Campos aninhados",
			Pessoa{
				"Maysa",
				Perfil{22, "Rio de Janeiro"},
			},
			[]string{"Maysa", "Rio de Janeiro"},
		},
		{
			"Ponteiros para coisas",
			&Pessoa{
				"Maysa",
				Perfil{22, "Rio de Janeiro"},
			},
			[]string{"Maysa", "Rio de Janeiro"},
		},
		{
			"Slices",
			[]Perfil{
				{20, "São Paulo"},
				{24, "Pará"},
			},
			[]string{"São Paulo", "Pará"},
		},
		{
			"Arrays",
			[2]Perfil{
				{20, "São Paulo"},
				{24, "Pará"},
			},
			[]string{"São Paulo", "Pará"},
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

	t.Run("com maps", func(t *testing.T) {
		mapA := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var resultado []string
		percorre(mapA, func(entrada string) {
			resultado = append(resultado, entrada)
		})

		verificaSeContem(t, resultado, "Bar")
		verificaSeContem(t, resultado, "Boz")
	})

}

type Pessoa struct {
	Nome   string
	Perfil Perfil
}

type Perfil struct {
	Idade  int
	Cidade string
}
