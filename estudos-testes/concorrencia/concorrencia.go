package concorrencia

import "time"

type VerificadorWebsite func(string) bool

func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			resultados[u] = vw(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return resultados
}
