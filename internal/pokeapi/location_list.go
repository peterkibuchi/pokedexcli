package pokeapi

import (
	"encoding/json"
	"log"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
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

	var locationsAreaResponse LocationAreasResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationsAreaResponse); err != nil {
		return LocationAreasResponse{}, err
	}

	return locationsAreaResponse, nil
}
