package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func newClient() Client {
	return Client {
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type LocationAreasResp struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}
