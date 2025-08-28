package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

func main() {
	// structs
	p := pessoa{"João", 30}
	fmt.Println(p)
	fmt.Println(p.nome)
	fmt.Println(p.idade)

	p.idade = 31
	fmt.Println(p.idade)

	p2 := pessoa{nome: "Maria"}
	fmt.Println(p2)
	fmt.Println(p2.nome)
	fmt.Println(p2.idade)

	p3 := &pessoa{nome: "Ana", idade: 25}
	fmt.Println(p3)
	fmt.Println(p3.nome)
	fmt.Println(p3.idade)

	p3.idade = 26
	fmt.Println(p3.idade)
}