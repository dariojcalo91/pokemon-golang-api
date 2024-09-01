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

type PokemonDetail struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Types []Type `json:"types"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonList struct {
	Results []Result `json:"results"`
}

type PokemonDTO struct {
	ID   int
	Name string
	Type string
}

func (pokemonDTO *PokemonDTO) Pokemon(ctx context.Context) (*pokemon.Pokemon, error) {
	pokemonParams := &pokemon.PokemonParams{
		ID:   pokemonDTO.ID,
		Name: pokemonDTO.Name,
		Type: pokemonDTO.Type,
	}

	p, err := pokemon.NewPokemon(pokemonParams)
	if err != nil {
		return nil, fmt.Errorf("%w with id %d: %w", ErrCreatingPokemon, pokemonDTO.ID, err)
	}

	return p, nil
}
