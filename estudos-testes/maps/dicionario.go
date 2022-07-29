package main

import "errors"

type Dicionario map[string]string

var (
	ErrNaoEncontrado    = errors.New("não foi possível encontar a palavra que você procura")
	ErrPalavraExistente = errors.New("não é possível adicionar a palavra pois ela já existe")
)

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]

	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, err := d.Busca(palavra)

	switch err {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrPalavraExistente
	default:
		return err
	}

	return nil
}
