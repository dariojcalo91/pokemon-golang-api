package pokerepo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dario-labs/srv/domain/pokemon"
	"github.com/dario-labs/srv/domain/specification"
	"github.com/dario-labs/srv/repository"
	"strconv"
)

const endpointPokemon = "https://pokeapi.co/api/v2/pokemon?limit=%s&offset=%s"

func (p *pokemonsRepository) List(ctx context.Context, spcn *specification.Specification) (pokemon.Pokemons, error) {
	pokemonsDTO, err := p.sendListRequest(ctx, spcn)
	if err != nil {
		return nil, err
	}

	var pokemonList pokemon.Pokemons
	for _, pkmDTO := range *pokemonsDTO {
		pkmn, err := pkmDTO.Pokemon(ctx)
		if err != nil {
			return nil, err
		}
		pokemonList = append(pokemonList, *pkmn)
	}

	return pokemonList, nil
}

func (p *pokemonsRepository) sendListRequest(ctx context.Context, spcn *specification.Specification) (*[]PokemonDTO, error) {
	endpoint := fmt.Sprintf(endpointPokemon, strconv.Itoa(spcn.PerPage), strconv.Itoa(spcn.Page))

	res, err := p.client.R().Get(endpoint)
	if err != nil {
		return nil, repository.HandleError(ctx, fmt.Sprintf("an error occurred calling to %s", endpoint), err)
	}
	if res.IsError() {
		return nil, pokemon.ErrPokemonNotFound
	}

	var rawListResponse PokemonList
	err = json.Unmarshal(res.Body(), &rawListResponse)
	if err != nil {
		return nil, repository.HandleError(ctx, "error unmarshalling body to pokemon list", err)
	}

	var listPokemonDTOs []PokemonDTO
	for _, pkm := range rawListResponse.Results {
		pkmResp, err := p.client.R().Get(pkm.Url)
		if err != nil {
			return nil, repository.HandleError(ctx, fmt.Sprintf("an error occurred calling to %s", pkmResp), err)
		}
		if pkmResp.IsError() {
			return nil, pokemon.ErrPokemonNotFound
		}

		var pokemonDetail PokemonDetail
		err = json.Unmarshal(pkmResp.Body(), &pokemonDetail)
		if err != nil {
			return nil, repository.HandleError(ctx, "error unmarshalling body to pokemon detail", err)
		}

		listPokemonDTOs = append(listPokemonDTOs, PokemonDTO{
			ID:   pokemonDetail.ID,
			Name: pokemonDetail.Name,
			Type: pokemonDetail.Types[0].Type.Name,
		})
	}

	return &listPokemonDTOs, nil
}
