package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type RequestConfig struct {
	nextLocationUrl     *string
	previousLocationUrl *string
}

func GetLocationRequest(url string) (LocationResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, err
	}

	var locationResponse LocationResponse
	if err := json.Unmarshal(data, &locationResponse); err != nil {
		return LocationResponse{}, err
	}

	return locationResponse, nil
}
