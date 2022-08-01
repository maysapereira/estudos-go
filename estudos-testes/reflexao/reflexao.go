package main

import "reflect"

func percorre(x interface{}, fn func(entrada string)) {
	valor := obtemValor(x)

	if valor.Kind() == reflect.Slice {
		for i := 0; i < valor.Len(); i++ {
			percorre(valor.Index(i).Interface(), fn)
		}

		return
	}

	for i := 0; i < valor.NumField(); i++ {
		campo := valor.Field(i)

		switch campo.Kind() {
		case reflect.String:
			fn(campo.String())
		case reflect.Struct:
			percorre(campo.Interface(), fn)
		}
	}
}

func obtemValor(x interface{}) reflect.Value {
	valor := reflect.ValueOf(x)

	if valor.Kind() == reflect.Ptr {
		valor = valor.Elem()
	}

	return valor
}
