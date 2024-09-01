package repositories

import (
	"context"
	"github.com/dario-labs/srv/domain/specification"

	"github.com/dario-labs/srv/domain/pokemon"
)

type PokemonRepository interface {
	Get(ctx context.Context, pokemonID int) (*pokemon.Pokemon, error)
	List(ctx context.Context, spec *specification.Specification) (pokemon.Pokemons, error)
}
