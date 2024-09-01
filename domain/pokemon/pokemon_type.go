package pokemon

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrInvalidPokemonType = errors.New("invalid pokemon type")

type knowPokemonType int

const (
	Normal knowPokemonType = iota
	Fire
	Water
	Electric
	Grass
	Ice
	Fighting
	Poison
	Ground
	Flying
	Psychic
	Bug
	Rock
	Ghost
	Dragon
	Dark
	Steel
	Fairy
)

type PokemonType struct {
	value knowPokemonType
}

func (p *PokemonType) Value() int {
	return int(p.value)
}

func (p *PokemonType) String() string {
	return strconv.Itoa(int(knowPokemonTypeMap[strconv.Itoa(int(p.value))]))
}

var (
	knowPokemonTypeMap = map[string]knowPokemonType{
		"normal": Normal,
		"fire":   Fire,
		"water":  Water,
		"grass":  Grass,
		"fairy":  Fairy,
		// TODO map all types
	}
)

func mapToPokemonType(input string) (knowPokemonType, bool) {
	c, ok := knowPokemonTypeMap[input]

	return c, ok
}

func NewPokemonType(value string) (*PokemonType, error) {
	pokemonType, ok := mapToPokemonType(value)
	if !ok {
		return nil, fmt.Errorf("%w, '%s' is not a valid type", ErrInvalidPokemonType, value)
	}

	return &PokemonType{
		value: pokemonType,
	}, nil
}
