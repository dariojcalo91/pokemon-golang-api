package shared

import (
	"cmp"

	"github.com/dario-labs/srv/domain/shared"
	sortingRules "github.com/dario-labs/srv/domain/sorter"
)

var orderCompareMap = map[sortingRules.Order]func(a, b int) int{
	sortingRules.AscendingOrder:  ascendingCompare,
	sortingRules.DescendingOrder: descendingCompare,
}

func ascendingCompare(a, b int) int {
	return cmp.Compare(a, b)
}

func descendingCompare(a, b int) int {
	return cmp.Compare(b, a)
}

const equal = 0

// SortingFn Función recursiva para ordenar items según []SortingRule.
// En caso de que al comparar 2 item el resultado es que son iguales
// recursivamente se compara con el siguiente SortingRule dentro del array
// en caso de que no queden SortingRules a comparar se considera que ambos item son iguales.
func SortingFn(sortingRules []sortingRules.SortingRule, sortingRulesIndex int, aMap, bMap map[shared.Field]int) int {
	if sortingRulesIndex == len(sortingRules) {
		return equal
	}
	s := sortingRules[sortingRulesIndex]
	orderFn := orderCompareMap[s.Order()]
	compareResult := orderFn(aMap[s.Field()], bMap[s.Field()])
	if compareResult == equal {
		return SortingFn(sortingRules, sortingRulesIndex+1, aMap, bMap)
	}
	return compareResult
}
