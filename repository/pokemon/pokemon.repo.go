package pokerepo

import (
	"github.com/dario-labs/srv/domain/repositories"
	restyV2 "github.com/go-resty/resty/v2"
)

type pokemonsRepository struct {
	client *restyV2.Client
}

func NewPokemonsRepository(client *restyV2.Client) repositories.PokemonRepository {
	return &pokemonsRepository{
		client: client,
	}
}
