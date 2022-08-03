package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	CPF       string
}

type PessoaJuridica struct {
	RazaoSocial string
	CNPJ        string
}

func (p PessoaFisica) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

func main() {
	p := PessoaFisica{
		Pessoa{Nome: "Maysa", Idade: 22, Status: true},
		"Pereira",
		"123.456.789-01",
	}

	fmt.Println(p)
}
