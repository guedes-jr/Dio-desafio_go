package main

import (
	"io"
	"os"
	"log"
)

func main() {
	if _, err := io.WriteString(os.Stdout, "Olá mundo!"); err != nil {
		log.Fatal(err)
	}
}