package pokeapi

type PokemonInfo struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

type PokeballInfo struct {
	Pokelist map[string]PokemonInfo
}
