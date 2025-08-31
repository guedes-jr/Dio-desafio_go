package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
)

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto")

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(responseData))
}
