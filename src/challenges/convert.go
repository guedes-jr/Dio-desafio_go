package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Digite o valor do ponto de ebulição da água em Kelvin: ")

    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Erro ao ler a entrada:", err)
        return
    }

    input = strings.TrimSpace(input)
    kelvin, err := strconv.ParseFloat(input, 64)
    if err != nil {
        fmt.Println("Por favor, digite um número válido.")
        return
    }

    if kelvin < 0 {
        fmt.Println("O valor em Kelvin não pode ser negativo.")
        return
    }

    celsius := kelvin - 273
    fmt.Printf("O valor de %.2fK corresponde a %.2f°C\n", kelvin, celsius)
}
