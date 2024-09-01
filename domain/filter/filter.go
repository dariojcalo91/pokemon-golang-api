package filter

import "github.com/dario-labs/srv/domain/shared"

const ( //other cases: emulate sql query filters
	FilterOperatorStringEqual = "string equal"
)

type Filter interface {
	Field() shared.Field
	Operator() string
	Value() string
}

type Attributes struct {
	field    shared.Field
	operator string
	value    string
}

func NewAttributes(field shared.Field, operator, value string) Attributes {
	return Attributes{
		field:    field,
		operator: operator,
		value:    value,
	}
}

func (a Attributes) Field() shared.Field {
	return a.field
}

func (a Attributes) Operator() string {
	return a.operator
}

func (a Attributes) Value() string {
	return a.value
}
