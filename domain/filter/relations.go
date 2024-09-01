package filter

type LogicalOperator string

const (
	AndOperator  LogicalOperator = "AND"
	OrOperator   LogicalOperator = "OR"
	SameOperator LogicalOperator = "SAME"
)

type Element interface {
	any
}

type LogicalRelationOfFilters LogicalRelation[Filter]

type LogicalRelation[T Element] interface {
	Elements() []T
	Operator() LogicalOperator
}

type relation[T Element] struct {
	elements []T
	operator LogicalOperator
}

func (r relation[T]) Elements() []T {
	return r.elements
}

func (r relation[T]) Operator() LogicalOperator {
	return r.operator
}

type AndRelation[T Element] struct {
	relation[T]
}

func NewAndRelation[T Element](elements ...T) LogicalRelation[T] {
	return AndRelation[T]{
		relation[T]{
			elements: elements,
			operator: AndOperator,
		},
	}
}

type SameRelation[T Element] struct {
	relation[T]
}

func NewSameRelation[T Element](element T) LogicalRelation[T] {
	return SameRelation[T]{
		relation[T]{
			elements: []T{element},
			operator: SameOperator,
		},
	}
}
