package main

import "errors"

type Dicionario map[string]string

var ErrNaoEncontrado = errors.New("não foi possível encontar a palavra que você procura")

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]

	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}
