package main
import "fmt"

type Pessoa struct {
    Nome string
    Idade int
}

func (p Pessoa) Apresentar() {
    fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}

func (p *Pessoa) Envelhecer() {
    p.Idade++
}

func main() {
    pessoa := Pessoa{"João", 30}
    pessoa.Apresentar()
    pessoa.Envelhecer()
    pessoa.Apresentar()
}
