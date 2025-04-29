package main

import (
	"fmt"
	"relational-algorithms/fdependency"
	"relational-algorithms/set"
)

func main() {

	attributes := set.NewFromElements([]string{"A", "B", "C", "D", "E", "F"})
	dep1 := fdependency.NewDepedency(*set.NewFromElements([]string{"B", "F"}), *set.NewFromElements([]string{"A", "C", "E"}))
	dep2 := fdependency.NewDepedency(*set.NewFromElements([]string{"A", "D", "F"}), *set.NewFromElements([]string{"B", "C", "E"}))
	dep3 := fdependency.NewDepedency(*set.NewFromElements([]string{"A", "B", "E"}), *set.NewFromElements([]string{"D"}))

	rel := fdependency.NewRelation(attributes,
		[]*fdependency.FunctionalDependency{dep1, dep2, dep3})

	rel.PerformSynthesisAlgorithm()

	fmt.Println(rel.String())

}
