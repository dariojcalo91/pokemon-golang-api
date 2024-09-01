package sorter

import "github.com/dario-labs/srv/domain/shared"

type Order string

const (
	AscendingOrder  Order = "ascending"
	DescendingOrder Order = "descending"
	CustomOrder     Order = "custom"
)

type SortingRule interface {
	Field() shared.Field
	Order() Order
}

type Sorter[T any] interface {
	Sort([]SortingRule, T)
}

type SortingRuleAttributes struct {
	field shared.Field
	order Order
}

func (s SortingRuleAttributes) Field() shared.Field {
	return s.field
}

func (s SortingRuleAttributes) Order() Order {
	return s.order
}
