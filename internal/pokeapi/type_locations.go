package pokeapi

type LocationMap struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous string          `json:"previous"`
	Results  []LocationNmUrl `json:"results"`
}

type LocationNmUrl struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
