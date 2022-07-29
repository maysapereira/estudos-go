package main

type Dicionario map[string]string

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

const (
	ErrNaoEncontrado      = ErrDicionario("não foi possível encontar a palavra que você procura")
	ErrPalavraExistente   = ErrDicionario("não é possível adicionar a palavra pois ela já existe")
	ErrPalavraInexistente = ErrDicionario("não foi possível atualizar a palavra pois ela não existe")
)

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

func (d Dicionario) Atualiza(palavra, definicao string) error {
	_, err := d.Busca(palavra)

	switch err {
	case ErrNaoEncontrado:
		return ErrPalavraInexistente
	case nil:
		d[palavra] = definicao
	default:
		return err
	}

	return nil
}

func (d Dicionario) Deleta(palavra string) {
	delete(d, palavra)
}
