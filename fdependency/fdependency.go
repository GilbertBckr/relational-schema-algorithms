package fdependency

import (
	"relational-algorithms/set"
)

type functionalDependency struct {
	Determinant set.Set
	Attributes  set.Set
}

func NewDepedency(determinant set.Set, attributes set.Set) *functionalDependency {
	return &functionalDependency{determinant, attributes}
}

type Relation struct {
	attributes set.Set
	// We make this a slice and not a set since the order can play a role for some algorithms
	functionalDependencies []*functionalDependency
}

func NewRelation(attributes *set.Set, deps []*functionalDependency) *Relation {
	return &Relation{
		attributes:             *attributes,
		functionalDependencies: deps,
	}

}

func Hull(r *Relation, determinants *set.Set) *set.Set {
	hull := determinants.DeepCopy()

	// basically graph search right?
	for {
		prevHull := hull.DeepCopy()
		for _, dep := range r.functionalDependencies {
			if dep.Determinant.IsSubSet(hull) {
				hull.AddUnion(&dep.Attributes)
			}

		}

		if prevHull.Equal(hull) {
			return hull
		}
	}

}
