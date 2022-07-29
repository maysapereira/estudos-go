package main

import (
	"fmt"
	"io"
	"os"
)

const ultimaPalavra = "Go!"
const inicioContagem = 3

func Contagem(saida io.Writer) {
	for i := inicioContagem; i > 0; i-- {
		fmt.Fprint(saida, i)
	}

	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	Contagem(os.Stdout)
}
