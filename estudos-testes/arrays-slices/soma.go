package main

func Soma(numeros []int) int {

	soma := 0

	for _, numero := range numeros {
		soma += numero
	}

	return soma
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {

	quantidadeDeNumeros := len(numerosParaSomar)
	somas = make([]int, quantidadeDeNumeros)

	for i, numeros := range numerosParaSomar {
		somas[i] = Soma(numeros)
	}

	return
}
