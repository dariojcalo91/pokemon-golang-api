package pokerepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/dario-labs/srv/domain/pokemon"
)

var ErrCreatingPokemon = errors.New("error creating pokemon from pokemonDTO")

type Type struct {
	Slot int
	Type struct {
		Name string `json:"name"`
	} `json:"Type"`
}

type Move struct {
	Move struct {
		Name string `json:"name"`
	} `json:"Move"`
}

type PokemonDetail struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Types  []Type `json:"types"`
	Moves  []Move `json:"moves"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonList struct {
	Results []Result `json:"results"`
}

type PokemonDTO struct {
	ID     int
	Name   string
	TypeA  string
	TypeB  string
	Moves  []string
	Height int
	Weight int
}

func (pokemonDTO *PokemonDTO) Pokemon(ctx context.Context) (*pokemon.Pokemon, error) {
	pokemonParams := &pokemon.PokemonParams{
		ID:     pokemonDTO.ID,
		Name:   pokemonDTO.Name,
		TypeA:  pokemonDTO.TypeA,
		TypeB:  pokemonDTO.TypeB,
		Moves:  pokemonDTO.Moves,
		Height: pokemonDTO.Height,
		Weight: pokemonDTO.Weight,
	}

	p, err := pokemon.NewPokemon(pokemonParams)
	if err != nil {
		return nil, fmt.Errorf("%w with id %d: %w", ErrCreatingPokemon, pokemonDTO.ID, err)
	}

	return p, nil
}
