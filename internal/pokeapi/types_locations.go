package pokeapi

// Generated at https://mholt.github.io/json-to-go/ with the endpoint output (gotten from the PokeAPI Docs)
type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
