package fdependency

import (
	"fmt"
	"reflect"
	"relational-algorithms/set"
	"strings"
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

func (r *Relation) Equals(r2 *Relation) bool {
	return r.attributes.Equals(&r2.attributes) && reflect.DeepEqual(r.functionalDependencies, r2.functionalDependencies)
}

func (r *Relation) String() string {
	builder := strings.Builder{}
	for _, i := range r.functionalDependencies {
		builder.WriteString(fmt.Sprintf("%v -> %v\n", i.Determinant.GetElementsOrdered(), i.Attributes.GetElementsOrdered()))
	}
	return builder.String()
}

func Hull(determinants *set.Set, functionalDependencies []*functionalDependency) *set.Set {
	hull := determinants.DeepCopy()

	// basically graph search right?
	for {
		prevHull := hull.DeepCopy()
		for _, dep := range functionalDependencies {
			if dep.Determinant.IsSubSet(hull) {
				hull.AddUnion(&dep.Attributes)
			}

		}

		if prevHull.Equals(hull) {
			return hull
		}
	}

}
