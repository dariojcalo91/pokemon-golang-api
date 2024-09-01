package shared

import (
	"errors"
	"fmt"
)

var ErrInvalidCriteria = errors.New("invalid criteria")

type CriteriaType string

const (
	Name   CriteriaType = "name"
	Order  CriteriaType = "order"
	Height CriteriaType = "height"
	//TODO add more criteria filters
)

var (
	knownCriteriaTypeMap = map[string]CriteriaType{
		"name":   Name,
		"order":  Order,
		"height": Height,
	}
)

func mapToCriteriaType(input string) (CriteriaType, bool) {
	c, ok := knownCriteriaTypeMap[input]
	return c, ok
}

type Criteria struct {
	value CriteriaType
}

func (c *Criteria) String() string {
	return string(c.value)
}

func (c *Criteria) Value() CriteriaType {
	return c.value
}

func NewCriteria(value string) (*Criteria, error) {
	criteria, ok := mapToCriteriaType(value)
	if !ok {
		return nil, fmt.Errorf("%w, '%s' is not a criteria", ErrInvalidCriteria, value)
	}
	return &Criteria{
		value: criteria,
	}, nil
}
