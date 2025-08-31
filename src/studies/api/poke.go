package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.get("http://pokeapi.co/api/v2/pokedex/kanto")

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.Read
}