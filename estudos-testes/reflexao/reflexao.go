package main

import (
	"reflect"
	"testing"
)

func percorre(x interface{}, fn func(entrada string)) {
	valor := obtemValor(x)

	quantidadeDeValores := 0
	var obtemCampo func(int) reflect.Value

	switch valor.Kind() {
	case reflect.String:
		fn(valor.String())
	case reflect.Struct:
		quantidadeDeValores = valor.NumField()
		obtemCampo = valor.Field
	case reflect.Slice, reflect.Array:
		quantidadeDeValores = valor.Len()
		obtemCampo = valor.Index
	case reflect.Map:
		for _, chave := range valor.MapKeys() {
			percorre(valor.MapIndex(chave).Interface(), fn)
		}
	}

	for i := 0; i < quantidadeDeValores; i++ {
		percorre(obtemCampo(i).Interface(), fn)
	}
}

func obtemValor(x interface{}) reflect.Value {
	valor := reflect.ValueOf(x)

	if valor.Kind() == reflect.Ptr {
		valor = valor.Elem()
	}

	return valor
}

func verificaSeContem(t *testing.T, palheiro []string, agulha string) {
	contem := false

	for _, x := range palheiro {
		if x == agulha {
			contem = true
		}
	}

	if !contem {
		t.Errorf("esperava-se que %+v contivesse '%s', mas não continha", palheiro, agulha)
	}
}
