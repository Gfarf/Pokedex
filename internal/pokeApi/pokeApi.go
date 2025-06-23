package pokeApi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}

// return explore
func (c *Client) ExploreLocations(loc string) (RespExploreLocations, error) {
	url := baseURL + "/location-area/" + loc
	if loc == "" {
		return RespExploreLocations{}, fmt.Errorf("Location was not provided.")
	}
	if val, ok := c.cache.Get(url); ok {
		exploreResp := RespExploreLocations{}
		err := json.Unmarshal(val, &exploreResp)
		if err != nil {
			return RespExploreLocations{}, err
		}

		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExploreLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExploreLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploreLocations{}, err
	}

	exploreResp := RespExploreLocations{}
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return RespExploreLocations{}, err
	}

	c.cache.Add(url, dat)
	return exploreResp, nil
}

// Return Pokemon stats
func (c *Client) ReturnPokemon(pokName string) (RespPokStats, error) {
	url := baseURL + "/pokemon/" + pokName
	if pokName == "" {
		return RespPokStats{}, fmt.Errorf("Pokemon was not provided.")
	}
	if val, ok := c.cache.Get(url); ok {
		pokeResp := RespPokStats{}
		err := json.Unmarshal(val, &pokeResp)
		if err != nil {
			return RespPokStats{}, err
		}

		return pokeResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokStats{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokStats{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokStats{}, err
	}

	pokeResp := RespPokStats{}
	err = json.Unmarshal(dat, &pokeResp)
	if err != nil {
		return RespPokStats{}, err
	}

	c.cache.Add(url, dat)
	return pokeResp, nil

}
