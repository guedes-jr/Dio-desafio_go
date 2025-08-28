package main

import "fmt"

func media(lista []float64) float64 {
	soma := 0.0
	for _, valor := range lista {
		soma += valor
	}
	return soma / float64(len(lista))
}

func main() {
    lista := []float64{94, 93, 77, 82, 83}

    fmt.Printf("Média = %.2f\n", media(lista))
}