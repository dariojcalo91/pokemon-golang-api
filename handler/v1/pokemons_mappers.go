package v1

import "github.com/dario-labs/srv/domain/pokemon"

type Move struct {
	Name string
}

type Type struct {
	Name string
}

type PokemonResponse struct {
	Id     int
	Name   string
	Types  []Type
	Moves  []Move
	Height int
	Weight int
}

func adaptPokemonModelToDTO(pokemon *pokemon.Pokemon) PokemonResponse {
	return PokemonResponse{
		Id:   pokemon.ID,
		Name: pokemon.Name,
		// Type: pokemon.Type.String(), //TODO fix wrong returned type data
	}
}
