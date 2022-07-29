package main

type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) string {
	return d[palavra]
}
