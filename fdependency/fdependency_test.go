package fdependency

import (
	"reflect"
	"relational-algorithms/set"
	"testing"
)

func TestHull(t *testing.T) {

	dep1 := NewDepedency(*set.NewFromElements([]string{"A", "B"}), *set.NewFromElements([]string{"C"}))
	dep2 := NewDepedency(*set.NewFromElements([]string{"D"}), *set.NewFromElements([]string{"E", "F"}))
	dep3 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"G", "H"}))
	dep4 := NewDepedency(*set.NewFromElements([]string{"G"}), *set.NewFromElements([]string{"B"}))

	rel := NewRelation(set.NewFromElements([]string{"A", "B", "C", "D", "E", "F", "G", "H"}),
		[]*FunctionalDependency{dep1, dep2, dep3, dep4})

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
		[]*FunctionalDependency{dep1, dep2, dep3, dep4, dep5, dep6})
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

	r := NewRelation(attributes, []*FunctionalDependency{dep1, dep2, dep3, dep4, dep5})

	r.leftReduction()

	reducedDep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "E"}))

	reducedDep2 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "D"}))
	reducedDep3 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D"}))
	reducedDep4 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"B", "E", "F"}))
	reducedDep5 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"B"}))

	exp := NewRelation(attributes, []*FunctionalDependency{reducedDep1, reducedDep2, reducedDep3, reducedDep4, reducedDep5})

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

	r := NewRelation(attributes, []*FunctionalDependency{dep1, dep2, dep3, dep4, dep5})

	r.rightReduction()

	reduceddep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"E"}))

	reduceddep2 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "D"}))
	reduceddep3 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D"}))
	reduceddep4 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"E", "F"}))
	reduceddep5 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"B"}))

	exp := NewRelation(attributes, []*FunctionalDependency{reduceddep1, reduceddep2, reduceddep3, reduceddep4, reduceddep5})

	if !exp.Equals(r) {
		t.Fatalf("Relation is not properly right reduced \n%s", r.String())
	}

}

func TestRemoveEmptyRules(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C", "D", "E", "F"})

	dep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "E"}))

	dep2 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "D"}))
	dep3 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D"}))
	dep4 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"B", "E", "F"}))
	dep5 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{}))

	r := NewRelation(attributes, []*FunctionalDependency{dep1, dep2, dep3, dep4, dep5})

	r.removeEmptyDependencies()

	reduceddep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "E"}))

	reduceddep2 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "D"}))
	reduceddep3 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D"}))
	reduceddep4 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"E", "F", "B"}))

	exp := NewRelation(attributes, []*FunctionalDependency{reduceddep1, reduceddep2, reduceddep3, reduceddep4})

	if !exp.Equals(r) {
		t.Fatalf("Empty rule was not removed \n%s", r.String())
	}

}

func TestMergeRule(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C", "D", "E", "F"})

	dep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "E"}))

	dep2 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "D"}))
	dep3 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D"}))
	dep4 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"B", "E", "F"}))
	dep5 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"B"}))

	r := NewRelation(attributes, []*FunctionalDependency{dep1, dep2, dep3, dep4, dep5})

	r.mergeRulesWithSameDeterminant()

	mergeddep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "D", "E"}))

	mergeddep2 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D", "B"}))
	mergeddep3 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"E", "F", "B"}))

	exp := NewRelation(attributes, []*FunctionalDependency{mergeddep1, mergeddep2, mergeddep3})

	if !exp.Equals(r) {
		t.Fatalf("rules were not properly merged \n%s", r.String())
	}

}

func TestCanonicalCover(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C", "D", "E", "F"})

	dep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "E"}))

	dep2 := NewDepedency(*set.NewFromElements([]string{"A", "E"}), *set.NewFromElements([]string{"B", "D"}))
	dep3 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D"}))
	dep4 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"B", "E", "F"}))
	dep5 := NewDepedency(*set.NewFromElements([]string{"C", "F"}), *set.NewFromElements([]string{"B"}))

	r := NewRelation(attributes, []*FunctionalDependency{dep1, dep2, dep3, dep4, dep5})

	r.CanonicalCover()

	mergeddep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "D", "E"}))

	mergeddep2 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D", "B"}))
	mergeddep3 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"E", "F"}))

	exp := NewRelation(attributes, []*FunctionalDependency{mergeddep1, mergeddep2, mergeddep3})

	if !exp.Equals(r) {
		t.Fatalf("canonical cover was not computed correctly \n%s", r.String())
	}

}
func TestCandidateKey(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C", "D", "E", "F"})

	dep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B", "E"}))

	dep2 := NewDepedency(*set.NewFromElements([]string{"A", "E"}), *set.NewFromElements([]string{"B", "D"}))
	dep3 := NewDepedency(*set.NewFromElements([]string{"F"}), *set.NewFromElements([]string{"C", "D"}))
	dep4 := NewDepedency(*set.NewFromElements([]string{"C", "D"}), *set.NewFromElements([]string{"B", "E", "F"}))
	dep5 := NewDepedency(*set.NewFromElements([]string{"C", "F"}), *set.NewFromElements([]string{"B"}))

	r := NewRelation(attributes, []*FunctionalDependency{dep1, dep2, dep3, dep4, dep5})

	candidateKeys := r.CandidateKeys()

	exp := []*set.Set{set.NewFromElements([]string{"A", "C"}), set.NewFromElements([]string{"A", "F"})}

	if !reflect.DeepEqual(candidateKeys, exp) {
		t.Fatalf("candidate key was not computer correctly \n%s", candidateKeys)
	}

}

func TestCandidateKey_SingleKey(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C"})

	dep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B"}))
	dep2 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"C"}))

	r := NewRelation(attributes, []*FunctionalDependency{dep1, dep2})

	candidateKeys := r.CandidateKeys()

	exp := []*set.Set{set.NewFromElements([]string{"A"})}

	if !reflect.DeepEqual(candidateKeys, exp) {
		t.Fatalf("candidate key was not computed correctly\n%s", candidateKeys)
	}
}

func TestCandidateKey_CompositeKey(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C", "D"})

	dep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B"}))
	dep2 := NewDepedency(*set.NewFromElements([]string{"B"}), *set.NewFromElements([]string{"C"}))

	r := NewRelation(attributes, []*FunctionalDependency{dep1, dep2})

	candidateKeys := r.CandidateKeys()

	exp := []*set.Set{set.NewFromElements([]string{"A", "D"})}

	if !reflect.DeepEqual(candidateKeys, exp) {
		t.Fatalf("candidate key was not computed correctly\n%s", candidateKeys)
	}
}
func TestCandidateKey_MultipleKeys(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C", "D"})

	dep1 := NewDepedency(*set.NewFromElements([]string{"A"}), *set.NewFromElements([]string{"B"}))
	dep2 := NewDepedency(*set.NewFromElements([]string{"C"}), *set.NewFromElements([]string{"D"}))

	r := NewRelation(attributes, []*FunctionalDependency{dep1, dep2})

	candidateKeys := r.CandidateKeys()

	exp := []*set.Set{
		set.NewFromElements([]string{"A", "C"}),
	}

	if !reflect.DeepEqual(candidateKeys, exp) {
		t.Fatalf("candidate key was not computed correctly\n%s", candidateKeys)
	}
}

func TestCandidateKey_NoFDs(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C"})

	r := NewRelation(attributes, []*FunctionalDependency{})

	candidateKeys := r.CandidateKeys()

	exp := []*set.Set{
		set.NewFromElements([]string{"A", "B", "C"}),
	}

	if !reflect.DeepEqual(candidateKeys, exp) {
		t.Fatalf("candidate key was not computed correctly\n%s", candidateKeys)
	}
}

func TestSynthesisAlgorithmSimple(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C", "D", "E"})
	dep1 := NewDepedency(*set.NewFromElements([]string{"B"}), *set.NewFromElements([]string{"A", "D", "E"}))
	dep2 := NewDepedency(*set.NewFromElements([]string{"B"}), *set.NewFromElements([]string{"C"}))
	dep3 := NewDepedency(*set.NewFromElements([]string{"A", "C", "E"}), *set.NewFromElements([]string{"B"}))

	rel := NewRelation(attributes, []*FunctionalDependency{dep1, dep2, dep3})

	//[B] -> [A B C D E]

	rel.PerformSynthesisAlgorithm()

	table1 := NewDepedency(*set.NewFromElements([]string{"B"}), *set.NewFromElements([]string{"A", "B", "C", "D", "E"}))

	if !rel.functionalDependencies[0].Equals(table1) {
		t.Fatalf("Does not match expected reduced schema %s", rel.String())
	}

}

// [B F] -> [A B E F]
//[A D F] -> [A B C D F]
//[A B E] -> [A B D E]

func TestSynthesisAlgorithm3TablesNoExtra(t *testing.T) {
	attributes := set.NewFromElements([]string{"A", "B", "C", "D", "E", "F"})
	dep1 := NewDepedency(*set.NewFromElements([]string{"B", "F"}), *set.NewFromElements([]string{"A", "C", "E"}))
	dep2 := NewDepedency(*set.NewFromElements([]string{"A", "D", "F"}), *set.NewFromElements([]string{"B", "C", "E"}))
	dep3 := NewDepedency(*set.NewFromElements([]string{"A", "B", "E"}), *set.NewFromElements([]string{"D"}))

	rel := NewRelation(attributes,
		[]*FunctionalDependency{dep1, dep2, dep3})

	rel.PerformSynthesisAlgorithm()

	table1 := NewDepedency(*set.NewFromElements([]string{"B", "F"}), *set.NewFromElements([]string{"A", "B", "E", "F"}))
	table2 := NewDepedency(*set.NewFromElements([]string{"A", "D", "F"}), *set.NewFromElements([]string{"A", "B", "C", "D", "F"}))
	table3 := NewDepedency(*set.NewFromElements([]string{"A", "B", "E"}), *set.NewFromElements([]string{"A", "B", "D", "E"}))

	if !rel.functionalDependencies[0].Equals(table1) || !rel.functionalDependencies[1].Equals(table2) || !rel.functionalDependencies[2].Equals(table3) {
		t.Fatalf("Does not match expected reduced schema %s", rel.String())

	}

}
