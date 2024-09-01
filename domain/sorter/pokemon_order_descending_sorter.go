package sorter

import "github.com/dario-labs/srv/domain/shared"

const (
	OrderField shared.Field = "order"
)

type ByOrderDescending struct {
	SortingRuleAttributes
}

func NewSorterByOrderDescending() ByOrderDescending {
	return ByOrderDescending{
		SortingRuleAttributes{
			field: OrderField,
			order: DescendingOrder,
		},
	}
}
