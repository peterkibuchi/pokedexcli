package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreasResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		fmt.Println("-> showing cached results...")
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err

	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	var locationsAreaResponse LocationAreasResponse
	err = json.Unmarshal(data, &locationsAreaResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(url, data)
	return locationsAreaResponse, nil
}
