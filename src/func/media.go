package main

import "fmt"

func main() {
    lista := []float64{94, 93, 77, 82, 83}
    soma := 0.0

    for _, valor := range lista {
        soma += valor
    }

    media := soma / float64(len(lista))
    fmt.Printf("Média = %.2f\n", media)
}