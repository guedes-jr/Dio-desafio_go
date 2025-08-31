package main

import "fmt"

func dia1 () {
	fmt.Println("Domingo")
} 

func dia2 () {
	fmt.Println("Segunda-feira")
}

func main() {
	defer dia1() // Adia a execução da função dia1 até o final da função main
	dia2()
}