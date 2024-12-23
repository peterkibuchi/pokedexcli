package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocationDetails(name *string) (LocationDetailsResponse, error) {
	url := baseUrl + "/location-area/" + *name

	if val, ok := c.cache.Get(url); ok {
		locationDetailsResp := LocationDetailsResponse{}
		err := json.Unmarshal(val, &locationDetailsResp)
		if err != nil {
			return LocationDetailsResponse{}, err
		}

		fmt.Println("-> showing cached results...")
		return locationDetailsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetailsResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetailsResponse{}, err

	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetailsResponse{}, err
	}

	var locationsAreaResponse LocationDetailsResponse
	err = json.Unmarshal(data, &locationsAreaResponse)
	if err != nil {
		return LocationDetailsResponse{}, err
	}

	c.cache.Add(url, data)
	return locationsAreaResponse, nil
}
