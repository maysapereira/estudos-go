package main

func Soma(numeros [5]int) int {

	soma := 0

	for i := 0; i < 5; i++ {
		soma += numeros[i]
	}

	return soma
}
