package specification

import (
	"github.com/dario-labs/srv/domain/filter"
	"github.com/dario-labs/srv/domain/sorter"
)

type Specification struct {
	LogicalRelations filter.LogicalRelation[filter.LogicalRelationOfFilters]
	Page             int
	PerPage          int
	SortingRules     []sorter.SortingRule
}

func NewSpecification(logicalRelations filter.LogicalRelation[filter.LogicalRelationOfFilters], page, perPage int, sorters []sorter.SortingRule) *Specification {
	return &Specification{LogicalRelations: logicalRelations, Page: page, PerPage: perPage, SortingRules: sorters}
}
