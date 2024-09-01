package v1

import (
	"context"
	"github.com/dario-labs/srv/domain/filter"
	"github.com/dario-labs/srv/domain/shared"
	"github.com/dario-labs/srv/usecases"
)

type PokemonService struct {
	UseCases *usecases.PokemonUseCases
}

type Params struct {
	Page     int
	PerPage  int
	Criteria string
}

func (p *PokemonService) GetPokemons(ctx context.Context, params Params) ([]PokemonResponse, error) {
	var response = []PokemonResponse{}

	filtersMap, err := buildFiltersMap(params)
	if err != nil {
		return response, err
	}

	pkmList, err := p.UseCases.List(ctx, filtersMap, params.Page, params.PerPage)
	if err != nil {
		return nil, err
	}

	for i := range pkmList {
		p := adaptPokemonModelToDTO(&pkmList[i])
		response = append(response, p)
	}

	return response, nil
}

func (p *PokemonService) GetPokemonDetail(ctx context.Context, pokemonID int) (PokemonResponse, error) {
	var response = PokemonResponse{}
	pkm, err := p.UseCases.Get(ctx, pokemonID)
	if err != nil {
		return response, err
	}
	response = adaptPokemonModelToDTO(pkm)

	return response, nil
}

func buildFiltersMap(params Params) (map[shared.FilterKey]filter.Filter, error) {
	return filter.NewFiltersMapBuilder().
		SetOptionalByCriteriaFilter(params.Criteria).
		Build()
}
