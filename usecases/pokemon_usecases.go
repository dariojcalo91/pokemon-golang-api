package usecases

import (
	"context"
	"github.com/dario-labs/srv/domain/filter"
	"github.com/dario-labs/srv/domain/specification"

	"github.com/dario-labs/srv/domain/pokemon"
	"github.com/dario-labs/srv/domain/repositories"
	"github.com/dario-labs/srv/domain/sorter"
)

type PokemonUseCases struct {
	repos         map[pokemon.PokemonGenIDPrefixType]repositories.PokemonRepository
	pokemonSorter sorter.Sorter[pokemon.Pokemons]
}

func NewPokemonUseCases(repos map[pokemon.PokemonGenIDPrefixType]repositories.PokemonRepository, pokemonSorter sorter.Sorter[pokemon.Pokemons]) *PokemonUseCases {
	return &PokemonUseCases{
		repos:         repos,
		pokemonSorter: pokemonSorter,
	}
}

func (p *PokemonUseCases) List(ctx context.Context, filtersMap filter.FiltersMap, page, perPage int) (pokemon.Pokemons, error) {
	spcn := specification.NewListPokemonSpecificationBasedOnCriteria(ctx, filtersMap, page, perPage)
	pokemons, err := p.repos[pokemon.Gen1].List(ctx, spcn)
	if err != nil {
		return nil, err
	}

	p.pokemonSorter.Sort(spcn.SortingRules, pokemons)

	return pokemons, nil
}

func (p *PokemonUseCases) Get(ctx context.Context, pokemonID int) (*pokemon.Pokemon, error) {
	return p.repos[pokemon.Gen1].Get(ctx, pokemonID)
}
