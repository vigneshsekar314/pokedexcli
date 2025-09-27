package pokeapi

func (c *Client) GetPokemonInfo(pokemonName string) (PokemonInfo, error) {
	url := BaseUrl + "pokemon/" + pokemonName
	resp, err := GetResponse[PokemonInfo](c, url, "Invalid pokemon name")
	if err != nil {
		return PokemonInfo{}, err
	}
	return resp, nil
}
