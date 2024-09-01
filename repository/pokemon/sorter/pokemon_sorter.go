package sorter

import (
	"slices"

	"github.com/dario-labs/srv/domain/pokemon"
	sortingRules "github.com/dario-labs/srv/domain/sorter"
	"github.com/dario-labs/srv/repository/shared"
)

type PokemonSorter struct{}

func NewPokemonSorter() *PokemonSorter {
	return &PokemonSorter{}
}

func (ps *PokemonSorter) Sort(rules []sortingRules.SortingRule, pokemons pokemon.Pokemons) {
	slices.SortFunc(pokemons,
		func(a, b pokemon.Pokemon) int {
			return shared.SortingFn(rules, 0, a.SortingRulesPrimitivesMap(), b.SortingRulesPrimitivesMap())
		})
}
