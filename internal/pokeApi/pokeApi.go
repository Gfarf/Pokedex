package pokeApi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Locs struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Named struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Locs `json:"results"`
}

func GetLocations(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return []string{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, res.Body)
	}
	if err != nil {
		return []string{}, err
	}

	data := Named{}
	if err := json.Unmarshal(body, &data); err != nil {
		return []string{}, err
	}
	var final []string
	for _, item := range data.Results {
		final = append(final, item.Name)
	}
	return final, nil
}

func GetNextPrevious(url string) (string, string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "null", "null", err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return "null", "null", fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, res.Body)
	}
	if err != nil {
		return "null", "null", err
	}

	data := Named{}
	if err := json.Unmarshal(body, &data); err != nil {
		return "null", "null", err
	}
	return data.Next, data.Previous, nil
}
