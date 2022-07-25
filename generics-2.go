package main

import "fmt"

func sum[V int64 | float64](m []V) V {
	var s V
	for _, v := range m {
		s += v

	}

	return s

	// [V int64 | float64] declara tipos de dados
	// m: valores recebidos em um slice chamado V
	// retorna a soma de V

	// var s do tipo V
	// intera por todos os valores passados em um for
	// dentro do for, s é somado com os valores de v
	// retorna valor de s
}

func main() {
	ints := []int64{1, 2, 3, 4, 5}
	floats := []float64{10.3, 3.7, 5.14, 3.14}

	fmt.Println("soma int64:", sum[int64](ints))
	fmt.Println("soma float64:", sum[float64](floats))
}
