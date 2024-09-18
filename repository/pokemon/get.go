package pokerepo

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/dario-labs/srv/domain/pokemon"
	"github.com/dario-labs/srv/repository"
)

const getPokemonEndpoint = "https://pokeapi.co/api/v2/pokemon/%s"

func (p *pokemonsRepository) Get(ctx context.Context, pokemonID int) (*pokemon.Pokemon, error) {
	pokemonDetailDTO, err := p.sendRequest(ctx, pokemonID)
	if err != nil {
		return nil, err
	}

	return pokemonDetailDTO.Pokemon(ctx)
}

func (p *pokemonsRepository) sendRequest(ctx context.Context, pokemonID int) (*PokemonDTO, error) {
	endpoint := fmt.Sprintf(getPokemonEndpoint, strconv.Itoa(pokemonID))

	res, err := p.client.R().Get(endpoint)
	if err != nil {
		return nil, repository.HandleError(ctx, fmt.Sprintf("an error occurred calling to %s", endpoint), err)
	}

	if res.IsError() {
		return nil, pokemon.ErrPokemonNotFound
	}

	var pokemonDetail PokemonDetail
	err = json.Unmarshal(res.Body(), &pokemonDetail)
	if err != nil {
		return nil, repository.HandleError(ctx, "error unmarshalling body to pokemon detail", err)
	}

	return &PokemonDTO{
		ID:    pokemonDetail.ID,
		Name:  pokemonDetail.Name,
		TypeA: pokemonDetail.Types[0].Type.Name,
	}, nil
}
