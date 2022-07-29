package main

import (
	"fmt"
	"io"
	"os"
)

func Contagem(saida io.Writer) {
	fmt.Fprint(saida, "3")
}

func main() {
	Contagem(os.Stdout)
}
