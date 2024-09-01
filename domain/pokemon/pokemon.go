package pokemon

import (
	"errors"

	"github.com/dario-labs/srv/domain/shared"
	"github.com/dario-labs/srv/domain/sorter"
)

var ErrPokemonNotFound = errors.New("Pokemon not found")

type Pokemon struct {
	ID   int
	Name string
	Type *PokemonType
}

func (p *Pokemon) SortingRulesPrimitivesMap() map[shared.Field]int {
	return map[shared.Field]int{
		sorter.NameTypeField: p.Type.Value(),
	}
}

type PokemonParams struct {
	ID   int
	Name string
	Type string
}

func NewPokemon(params *PokemonParams) (*Pokemon, error) {
	pokemonType, err := NewPokemonType(params.Type)
	if err != nil {
		return nil, err
	}

	return &Pokemon{
		ID:   params.ID,
		Name: params.Name,
		Type: pokemonType,
	}, nil
}
