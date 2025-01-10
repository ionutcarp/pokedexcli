package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	if pokemonName == "" {
		return Pokemon{}, errors.New("Pokemon name is required")
	}
	fullURL := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(fullURL); ok {
		pokemonDetailFromCache := Pokemon{}
		err := json.Unmarshal(val, &pokemonDetailFromCache)
		if err != nil {
			return Pokemon{}, err
		}
		fmt.Println("Pokemon retrieved from cache")
		return pokemonDetailFromCache, nil
	}
	//fmt.Println("Retrieving Pokemon detail from HTTP")

	request, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	if response.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}

	PokemonDetailResponse := Pokemon{}
	err = json.Unmarshal(data, &PokemonDetailResponse)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(fullURL, data)
	return PokemonDetailResponse, nil
}
