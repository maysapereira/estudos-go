package main

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

func Area(retangulo Retangulo) float64 {
	return retangulo.Largura + retangulo.Altura

}

type Retangulo struct {
	Largura float64
	Altura  float64
}
