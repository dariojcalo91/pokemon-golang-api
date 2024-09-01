package v1

import "github.com/dario-labs/srv/domain/pokemon"

type PokemonResponse struct {
	Id   int
	Name string
	Type string
}

func adaptPokemonModelToDTO(pokemon *pokemon.Pokemon) PokemonResponse {
	return PokemonResponse{
		Id:   pokemon.ID,
		Name: pokemon.Name,
		Type: pokemon.Type.String(), //TODO fix wrong returned type data
	}
}
