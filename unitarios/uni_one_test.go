package unitarios_test

import (
	"reflect"
	"testing"

	"github.com/eron97/testes-go.git/unitarios"
)

/*

func TestSoma(t *testing.T) {
	resultado := unitarios.Soma(10, 20)

	if resultado != 30 {
		t.Errorf("Func Soma() retornou %d, esperado %d", resultado, 30)
	}
}

*/

/*
func TestSoma(t *testing.T) {
	t.Run("Teste Básico", func(t *testing.T) {
		t.Helper()
		resultado := unitarios.Soma(10, 20)
		if resultado != 30 {
			t.Errorf("retornou %d, esperado %d", resultado, 30)
		}
	})

	t.Run("Teste com Números Negativos", func(t *testing.T) {
		t.Helper()
		resultado := unitarios.Soma(-5, -10)
		if resultado != -10 {
			t.Errorf("retornou %d, esperado %d", resultado, -10)
		}
	})

	t.Run("Teste com Números Grandes", func(t *testing.T) {
		t.Helper()
		resultado := unitarios.Soma(1000, 5000)
		if resultado != 6000 {
			t.Errorf("retornou %d, esperado %d", resultado, 6000)
		}
	})
}
*/

/*
func TestSoma(t *testing.T) {

	testes := []struct {
		a, b, esperado int
		name_test      string
	}{
		{1, 2, 3, "inteiros positivos"},
		{-1, 1, 0, "resultado zero"},
		{0, 0, 0, "zeros"},
		{10, -15, -30, "resultado negativo"},
	}

	for _, teste := range testes {
		resultado := unitarios.Soma(teste.a, teste.b)
		if resultado != teste.esperado {
			t.Errorf("%s: esperado %d, retornou %d", teste.name_test, teste.esperado, resultado)
		}
	}

}
*/

/*
func TestCalcularMedia(t *testing.T) {
	testes := []struct {
		numeros   []int
		result    int
		name_test string
	}{
		{[]int{}, 0, "len de zero (slice vazio)"},
		{[]int{1, 2, 3}, 2, "positivos pequenos (soma = 6, média = 2)"},
		{[]int{100, 200, 300}, 200, "positivos grandes (soma = 600, média = 200)"},
		{[]int{0, 0, 0}, 0, "zeros (soma = 0, média = 0)"},
		{[]int{-10, -20, -30}, -20, "negativos (soma = -60, média = -20)"},
		{[]int{-20, 20, 0}, 0, "inteiros + e - (soma = 0, média = 0)"},
	}

	for _, teste := range testes {
		resultado := unitarios.CalcularMedia(teste.numeros)
		if resultado != float64(teste.result) {
			t.Errorf("%s: esperado %f, resultado %f", teste.name_test, float64(teste.result), resultado)
		}
	}
}
*/

func TestContarOcorrencias(t *testing.T) {
	testes := []struct {
		lista     []string
		result    map[string]int
		name_test string
	}{
		{[]string{"banana", "maça", "melancia"}, map[string]int{"banana": 1, "maça": 1, "melancia": 1}, "não repetidos"},
		{[]string{"banana", "uva", "banana", "melancia"}, map[string]int{"banana": 2, "uva": 1, "melancia": 1}, "repetidos"},
		{[]string{}, map[string]int{}, "entrada vazia"},
	}

	for _, teste := range testes {
		resultado := unitarios.ContarOcorrencias(teste.lista)

		if !reflect.DeepEqual(resultado, teste.result) {
			t.Errorf("%s: esperado %v, resultado %v", teste.name_test, teste.result, resultado)
		}
	}
}
