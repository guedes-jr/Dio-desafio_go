package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
)

type Response struct {
	name    string    `json: "name`
	pokemon []pokemon `json: "pokemon_entries`
}

type pokemon struct {
	numero  int            `json: "entry_number"`
	especie pokemonSpecies `json: "pokemon_species"`
}

type pokemonSpecies struct {
	name string `json: "name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto")

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Print(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Print(responseObject.name)
	fmt.Print((responseObject.pokemon))

	for _, pokemon := range responseObject.pokemon {
		fmt.Print(pokemon.especie.name)
	}
}
