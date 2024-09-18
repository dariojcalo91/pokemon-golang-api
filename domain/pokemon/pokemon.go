package pokemon

import (
	"errors"

	"github.com/dario-labs/srv/domain/shared"
	"github.com/dario-labs/srv/domain/sorter"
)

var ErrPokemonNotFound = errors.New("Pokemon not found")

type Pokemon struct {
	ID     int
	Name   string
	TypeA  *PokemonType
	TypeB  *PokemonType
	Moves  []string
	Height int
	Weight int
}

func (p *Pokemon) SortingRulesPrimitivesMap() map[shared.Field]int {
	return map[shared.Field]int{
		sorter.NameTypeField: p.TypeA.Value(),
	}
}

type PokemonParams struct {
	ID     int
	Name   string
	TypeA  string
	TypeB  string
	Moves  []string
	Height int
	Weight int
}

func NewPokemon(params *PokemonParams) (*Pokemon, error) {
	pokemonTypeA, err := NewPokemonType(params.TypeA)
	if err != nil {
		return nil, err
	}

	var pokemonTypeB *PokemonType
	if params.TypeB != "" {
		pokemonTypeB, err = NewPokemonType(params.TypeB)
		if err != nil {
			return nil, err
		}
	}

	return &Pokemon{
		ID:     params.ID,
		Name:   params.Name,
		TypeA:  pokemonTypeA,
		TypeB:  pokemonTypeB,
		Moves:  params.Moves,
		Height: params.Height,
		Weight: params.Weight,
	}, nil
}
