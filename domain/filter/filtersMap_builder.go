package filter

import "github.com/dario-labs/srv/domain/shared"

type FiltersMap map[shared.FilterKey]Filter

type FiltersMapBuilder struct {
	filtersMap FiltersMap
	err        error
}

func NewFiltersMapBuilder() *FiltersMapBuilder {
	return &FiltersMapBuilder{
		filtersMap: map[shared.FilterKey]Filter{},
	}
}

func (f *FiltersMapBuilder) SetByCriteriaFilter(criteria string) *FiltersMapBuilder {
	if f.err != nil {
		return f
	}
	f.filtersMap[CriteriaFilter], f.err = NewByCriteriaFilter(criteria)
	return f
}

func (f *FiltersMapBuilder) SetOptionalByCriteriaFilter(criteria string) *FiltersMapBuilder {
	if criteria == "" {
		return f
	}
	return f.SetByCriteriaFilter(criteria)
}

func (f *FiltersMapBuilder) Build() (map[shared.FilterKey]Filter, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.filtersMap, nil
}
