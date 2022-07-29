package main

import "errors"

type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]

	if !existe {
		return "", errors.New("não foi possível encontrar a palavra que você procura")
	}

	return definicao, nil
}
