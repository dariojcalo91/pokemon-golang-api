package filter

import "github.com/dario-labs/srv/domain/shared"

const (
	CriteriaField  shared.Field     = "criteria"
	CriteriaFilter shared.FilterKey = "criteria_filter"
)

type ByCriteria struct {
	Attributes
}

func NewByCriteriaFilter(criteria string) (ByCriteria, error) {
	newCriteria, err := shared.NewCriteria(criteria)
	if err != nil {
		return ByCriteria{}, err
	}
	return ByCriteria{
		Attributes: NewAttributes(
			CriteriaField,
			FilterOperatorStringEqual,
			newCriteria.String(),
		),
	}, nil
}
