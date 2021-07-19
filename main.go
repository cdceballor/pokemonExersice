package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-filter-pokemon-api/requests"
	"github.com/go-filter-pokemon-api/services"
)

func main() {
	//http://18.119.14.15:3000/api/v2/pokemon  - https://pokeapi.co/api/v2/pokemon/?offset=0&limit=1118
	os.Setenv("PokemonURL", "http://18.119.14.15:3000/api/v2/pokemon")
	//os.Setenv("PokemonURL", "https://pokeapi.co/api/v2/pokemon/?offset=0&limit=1118")

	/*
		r := gin.Default()

		controllers.InitFilterController(
			services.Filters{
				ApiRequest: &requests.PokeApiRequest{},
			},
			r,
		)

		r.Run()

	*/
	start := time.Now()
	a := services.Filters{
		ApiRequest: &requests.PokeApiRequest{},
	}
	v, _, _, _ := a.WeightAndHeight(500, 500)
	log.Printf("Time %s", time.Since(start))

	fmt.Println(v)
}
