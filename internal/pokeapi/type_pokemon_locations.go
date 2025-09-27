package pokeapi

type PokemonsByLocation struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemn Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
