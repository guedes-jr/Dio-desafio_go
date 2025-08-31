package main

import (
	"fmt"
	"time"
)

func ping(pingChan, pongChan chan bool) {
	for {
		<-pongChan // espera o "sinal" do pong
		fmt.Println("ping")
		time.Sleep(500 * time.Millisecond)
		pingChan <- true // libera o pong
	}
}

func pong(pingChan, pongChan chan bool) {
	for {
		<-pingChan // espera o "sinal" do ping
		fmt.Println("pong")
		time.Sleep(500 * time.Millisecond)
		pongChan <- true // libera o ping
	}
}

func main() {
	pingChan := make(chan bool)
	pongChan := make(chan bool)

	go ping(pingChan, pongChan)
	go pong(pingChan, pongChan)

	// Inicia o jogo liberando o primeiro "pong"
	pongChan <- true

	// Mantém o programa rodando
	select {}
}
