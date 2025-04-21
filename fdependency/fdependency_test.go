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

	hull := Hull(set.NewFromElements([]string{"A"}), rel.functionalDependencies)

	expected := set.NewFromElements([]string{"A", "B", "C", "G", "H"})

	if !hull.Equals(expected) {
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
	hull := Hull(set.NewFromElements([]string{"A"}), rel.functionalDependencies)
	expected := set.NewFromElements([]string{"A", "B", "C", "D", "E", "F", "G", "H"})

	if !hull.Equals(expected) {
		t.Fatalf("Sets do not match hull=%v \n expected=%v", hull, expected)
	}

}

func TestLeftReduction(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C", "D", "E", "F"})

	dep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "E"}))

	dep2 := NewDepedency(*set.NewFromElements([]string{"A", "E"}), *set.NewFromElements([]string{"B", "D"}))
	dep3 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D"}))
	dep4 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"B", "E", "F"}))
	dep5 := NewDepedency(*set.NewFromElements([]string{"C", "F"}), *set.NewFromElements([]string{"B"}))

	r := NewRelation(attributes, []*functionalDependency{dep1, dep2, dep3, dep4, dep5})

	r.LeftReduction()

	reducedDep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "E"}))

	reducedDep2 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "D"}))
	reducedDep3 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D"}))
	reducedDep4 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"B", "E", "F"}))
	reducedDep5 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"B"}))

	exp := NewRelation(attributes, []*functionalDependency{reducedDep1, reducedDep2, reducedDep3, reducedDep4, reducedDep5})

	if !exp.Equals(r) {
		t.Fatalf("Relation is not properly left reduced")
	}

}

func TestRightReduction(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C", "D", "E", "F"})

	dep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "E"}))

	dep2 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "D"}))
	dep3 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D"}))
	dep4 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"B", "E", "F"}))
	dep5 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"B"}))

	r := NewRelation(attributes, []*functionalDependency{dep1, dep2, dep3, dep4, dep5})

	r.rightReduction()

	reduceddep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"E"}))

	reduceddep2 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "D"}))
	reduceddep3 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D"}))
	reduceddep4 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"E", "F"}))
	reduceddep5 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"B"}))

	exp := NewRelation(attributes, []*functionalDependency{reduceddep1, reduceddep2, reduceddep3, reduceddep4, reduceddep5})

	if !exp.Equals(r) {
		t.Fatalf("Relation is not properly right reduced \n%s", r.String())
	}

}
