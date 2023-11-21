package unitarios

// FuncaoSimples retorna a soma de dois inteiros.
func Soma(a, b int) int {
	return a + b
}

// CalcularMedia calcula a média de uma lista de números.
func CalcularMedia(numeros []int) float64 {
	if len(numeros) == 0 {
		return 0
	}

	soma := 0
	for _, num := range numeros {
		soma += num
	}

	media := float64(soma) / float64(len(numeros))
	return media
}

// ContarOcorrencias conta a ocorrência de cada elemento em um slice de strings.
func ContarOcorrencias(lista []string) map[string]int {
	ocorrencias := make(map[string]int)

	for _, elemento := range lista {
		ocorrencias[elemento]++
	}

	return ocorrencias
}
