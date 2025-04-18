package fdependency

import (
	"relational-algorithms/set"
	"testing"
)

func TestHull(t *testing.T) {

	dep1 := NewDepedency(*set.NewFromElements([]string{"A", "B"}), *set.NewFromElements([]string{"C"}))
	dep2 := NewDepedency(*set.NewFromElements([]string{"D"}), *set.NewFromElements([]string{"E", "F"}))
	dep3 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"G", "H"}))
	dep4 := NewDepedency(*set.NewFromElements([]string{"G"}), *set.NewFromElements([]string{"B"}))

	rel := NewRelation(set.NewFromElements([]string{"A", "B", "C", "D", "E", "F", "G", "H"}),
		[]*functionalDependency{dep1, dep2, dep3, dep4})

	hull := Hull(rel, set.NewFromElements([]string{"A"}))

	expected := set.NewFromElements([]string{"A", "B", "C", "G", "H"})

	if !hull.Equal(expected) {
		t.Fatalf("Sets do not match hull=%v \n expected=%v", hull, expected)
	}

}

func TestHullFull(t *testing.T) {

	dep1 := NewDepedency(*set.NewFromElements([]string{"A", "B"}), *set.NewFromElements([]string{"C"}))
	dep2 := NewDepedency(*set.NewFromElements([]string{"D"}), *set.NewFromElements([]string{"E", "F"}))
	dep3 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"G", "H"}))
	dep6 := NewDepedency(*set.NewFromElements([]string{"D", "E"}), *set.NewFromElements([]string{"F", "G"}))
	dep4 := NewDepedency(*set.NewFromElements([]string{"G"}), *set.NewFromElements([]string{"B"}))
	dep5 := NewDepedency(*set.NewFromElements([]string{"G"}), *set.NewFromElements([]string{"D", "E"}))

	rel := NewRelation(set.NewFromElements([]string{"A", "B", "C", "D", "E", "F", "G", "H"}),
		[]*functionalDependency{dep1, dep2, dep3, dep4, dep5, dep6})

	hull := Hull(rel, set.NewFromElements([]string{"A"}))

	expected := set.NewFromElements([]string{"A", "B", "C", "D", "E", "F", "G", "H"})

	if !hull.Equal(expected) {
		t.Fatalf("Sets do not match hull=%v \n expected=%v", hull, expected)
	}

}
