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

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// RespExploreLocations
type RespExploreLocations struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	Pokemon_encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int `json:"min_level"`
				MaxLevel        int `json:"max_level"`
				ConditionValues []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"condition_values"`
				Chance int `json:"chance"`
				Method struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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
