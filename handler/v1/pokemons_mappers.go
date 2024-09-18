package v1

import "github.com/dario-labs/srv/domain/pokemon"

type PokemonResponse struct {
	Id     int
	Name   string
	TypeA  string
	TypeB  string
	Moves  []string
	Height int
	Weight int
}

func adaptPokemonModelToDTO(pokemon *pokemon.Pokemon) PokemonResponse {
	var typeB string
	if pokemon.TypeB != nil {
		typeB = pokemon.TypeB.String()
	}

	return PokemonResponse{
		Id:     pokemon.ID,
		Name:   pokemon.Name,
		TypeA:  pokemon.TypeA.String(),
		TypeB:  typeB,
		Moves:  pokemon.Moves,
		Height: pokemon.Height,
		Weight: pokemon.Weight,
	}
}
