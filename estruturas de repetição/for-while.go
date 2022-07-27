package forwhile

import "fmt"

func main() {

	frutas := "van escolar"

	switch frutas {

	case "uva", "banana":
		fmt.Println("CAIU NA PRIMEIRA CONDIÇÃO")

	case "morango", "amora":
		fmt.Println("CAIU NA SEGUNDA CONDIÇÃO")
	case "tangerina", "maçã":
		fmt.Println("CAIU NA TERCEIRA CONDIÇÃO")
	default:
		fmt.Println("Não é fruta")
	valor := 0

	for valor <= 10 {
		fmt.Println("VALOR:", valor)
		valor++
	}
}
}

// "while"