package sorter

import "github.com/dario-labs/srv/domain/shared"

const (
	NameTypeField shared.Field = "name"
)

type ByNameAscending struct {
	SortingRuleAttributes
}

func NewSorterByNameAscending() ByNameAscending {
	return ByNameAscending{
		SortingRuleAttributes{
			field: NameTypeField,
			order: AscendingOrder,
		},
	}
}
