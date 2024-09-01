package specification

import (
	"context"
	"github.com/dario-labs/srv/domain/filter"
	"github.com/dario-labs/srv/domain/shared"
	"github.com/dario-labs/srv/domain/sorter"
	"log"
)

var pokemonSortingRulesMap = map[shared.CriteriaType]func() []sorter.SortingRule{
	shared.Name:   NewByNamePokemonSortingRule,
	shared.Height: NewByHeightPokemonSortingRule,
}

func NewListPokemonSpecificationBasedOnCriteria(ctx context.Context, filtersMap filter.FiltersMap, page, perPage int) *Specification {
	pokemonCriteriaFilter, ok := filtersMap[filter.CriteriaFilter]
	if !ok {
		return NewPokemonSpecification(filtersMap, page, perPage, NewPokemonDefaultSortingRules())
	}
	sortingRules, ok := pokemonSortingRulesMap[shared.CriteriaType(pokemonCriteriaFilter.Value())]
	if !ok {
		log.Fatalf("unkown sorting rule: %s, applying default sorting rules \n", pokemonCriteriaFilter.Value())
		sortingRules = NewPokemonDefaultSortingRules
	}
	return NewPokemonSpecification(filtersMap, page, perPage, sortingRules())
}

func NewPokemonSpecification(filtersMap filter.FiltersMap, page, perPage int, sorters []sorter.SortingRule) *Specification {
	return NewSpecification(buildLogicalRelationsWithFilters(filtersMap), page, perPage, sorters)
}

func buildLogicalRelationsWithFilters(filtersMap filter.FiltersMap) filter.LogicalRelation[filter.LogicalRelationOfFilters] {
	var filters []filter.Filter
	for _, v := range filtersMap {
		filters = append(filters, v)
	}
	filtersAndRelation := filter.NewAndRelation[filter.Filter](filters...)
	return filter.NewSameRelation[filter.LogicalRelationOfFilters](filtersAndRelation)
}

func NewPokemonDefaultSortingRules() []sorter.SortingRule {
	return []sorter.SortingRule{sorter.NewSorterByOrderDescending()}
}

func NewByNamePokemonSortingRule() []sorter.SortingRule {
	return []sorter.SortingRule{sorter.NewSorterByNameAscending(), sorter.NewSorterByOrderDescending()}
}

func NewByHeightPokemonSortingRule() []sorter.SortingRule {
	return []sorter.SortingRule{sorter.NewSorterByHeightAscending()}
}
