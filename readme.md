# Projeto Go - Estudos e Exercícios

Este repositório contém diversos exemplos, exercícios e experimentos em Go, organizados em pastas temáticas. O objetivo é servir como material de estudo e referência para conceitos fundamentais da linguagem Go, incluindo arrays, slices, mapas, estruturas, métodos, controle de fluxo, conversão de tipos e manipulação de variáveis.

## Estrutura do Projeto

- **hello/**  
  Exemplos básicos de Go, incluindo manipulação de variáveis em [`vars.go`](hello/vars.go) e outros conceitos introdutórios em [`numero.go`](hello/numero.go).

- **pkg/**  
  Pasta reservada para pacotes reutilizáveis (atualmente vazia).

- **src/temp/**  
  Coleção de exemplos práticos:
  - [`array.go`](src/temp/array.go): Demonstração de arrays em Go.
  - [`fatia.go`](src/temp/fatia.go): Manipulação de slices, cópia e efeitos colaterais.
  - [`mapa.go`](src/temp/mapa.go): Uso de mapas (dicionários).
  - [`estrutura.go`](src/temp/estrutura.go): Definição e uso de structs.
  - [`metodo.go`](src/temp/metodo.go): Implementação de métodos em structs.
  - [`parImpar.go`](src/temp/parImpar.go): Verificação de números pares e ímpares.
  - [`switch.go`](src/temp/switch.go): Uso do comando switch.

- **src/challenges/**  
  Coleção de exemplos práticos:
  - [`convert.go`](src/challenges/convert.go): Conversão de tipos.
  - [`desafio_1.go`](src/challenges/desafio_1.go)
  - [`desafio_2.go`](src/challenges/desafio_2.go): Desafios práticos de lógica e programação.

- **src/func/**  
  Coleção de exemplos práticos:
  - [`media.go`](src/func/media.go): Função para calculo da mpedia de uma sala.

## Como Executar

Cada arquivo Go pode ser executado individualmente. Por exemplo, para rodar o exemplo de slices:

```sh
go run [fatia.go](http://_vscodecontentref_/0)