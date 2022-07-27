package main

func Soma(numeros [5]int) int {

	soma := 0

	for _, numero := range numeros {
		soma += numero
	}

	return soma
}
