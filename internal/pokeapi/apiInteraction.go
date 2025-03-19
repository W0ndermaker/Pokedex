package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Get Locations
func (c *Client) GetLocationArea(url string) (LocationAreaAPIResponse, error) {

	// checking cache
	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationAreaAPIResponse{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationAreaAPIResponse{}, err
		}
		return locationResp, nil
	}

	// Get request to Pokedex
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaAPIResponse{}, err
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationAreaAPIResponse{}, err
	}
	defer response.Body.Close()

	dat, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationAreaAPIResponse{}, err
	}

	var locationResponse LocationAreaAPIResponse
	err = json.Unmarshal(dat, &locationResponse)
	if err != nil {
		return LocationAreaAPIResponse{}, err
	}

	c.cache.Add(url, dat)
	return locationResponse, nil
}

// Explore location
func (c *Client) ExploreArea(url string) (Location, error) {
	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return Location{}, err
	}
	defer response.Body.Close()

	dat, err := io.ReadAll(response.Body)
	if err != nil {
		return Location{}, err
	}

	var PokemonResponse Location
	err = json.Unmarshal(dat, &PokemonResponse)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)
	return PokemonResponse, nil

}

func (c *Client) GetPokemon(url string) (Pokemon, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}
	defer response.Body.Close()

	dat, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var PokemonInfo Pokemon
	err = json.Unmarshal(dat, &PokemonInfo)
	if err != nil {
		return Pokemon{}, err
	}

	return PokemonInfo, nil

}
