package services

import (
	"fmt"
	"sync"

	"github.com/go-filter-pokemon-api/models"
	"github.com/go-filter-pokemon-api/requests"
)

type PokemonFilters interface {
	WeightAndHeight(height int, weight int) ([]*models.Pokemon, int, []error, error)
}

type Filters struct {
	ApiRequest requests.PokemonRequest
}

func (filters *Filters) TraerPokemon(c1 chan *models.Pokemon, c2 chan error, height int, weight int, pokemon *models.PokemonList) {
	wg := sync.WaitGroup{}
	for _, result := range pokemon.Results {
		wg.Add(1)
		go func(url string) {
			p, err := filters.ApiRequest.GetPokemonByUrlId(url)
			if err != nil {
				c2 <- err
			} else if p.Height >= height && p.Weight >= weight {
				c1 <- p
			}
			wg.Done()
		}(result.Url)
	}
	wg.Wait()
	close(c1)
	close(c2)
}

func (filters *Filters) WeightAndHeight(height int, weight int) ([]*models.Pokemon, int, []error, error) {
	var wg sync.WaitGroup
	defer wg.Done()
	wg.Add(1)
	pokemons, err := filters.ApiRequest.GetAllPokemon()
	pokemonChannel := make(chan *models.Pokemon, pokemons.Count)
	errChannel := make(chan error, pokemons.Count)

	if err != nil {
		return nil, 0, nil, err
	}
	go filters.TraerPokemon(pokemonChannel, errChannel, height, weight, pokemons)

	arr := []*models.Pokemon{}
	errs := []error{}
	for pokemons := range pokemonChannel {
		arr = append(arr, pokemons)
	}
	//Abrir dos go rutinas diferentes, una para pokemones y otra para errores
	for er := range errChannel {
		errs = append(errs, er)
	}

	if len(errs) > 0 {
		fmt.Println("**********************************************************")
		Err := fmt.Errorf("TOTAL DE POKÉMONES CON ERROR AL INGRESAR: ", len(errs))
		fmt.Println(Err)
		fmt.Println("**********************************************************")
		for problem := range errs {
			fmt.Println("Error #:", problem, "->", errs[problem])
		}
		//return arr, len(arr), errs, errors.New("Error en los datos")
	}

	return arr, len(arr), errs, nil
}
