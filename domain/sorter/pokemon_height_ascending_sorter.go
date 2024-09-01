package sorter

import "github.com/dario-labs/srv/domain/shared"

const (
	HeightField shared.Field = "height"
)

type ByHeightAscending struct {
	SortingRuleAttributes
}

func NewSorterByHeightAscending() ByHeightAscending {
	return ByHeightAscending{
		SortingRuleAttributes{
			field: HeightField,
			order: AscendingOrder,
		},
	}
}
