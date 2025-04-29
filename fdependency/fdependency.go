package fdependency

import (
	"fmt"
	"reflect"
	"relational-algorithms/set"
	"slices"
	"strings"
)

type FunctionalDependency struct {
	Determinant set.Set
	Attributes  set.Set
}

func (f *FunctionalDependency) getFullSet() *set.Set {
	res := f.Determinant.DeepCopy()
	res.AddUnion(&f.Attributes)
	return res
}

func (f *FunctionalDependency) Equals(f2 *FunctionalDependency) bool {
	return f.Attributes.Equals(&f2.Attributes) && f.Determinant.Equals(&f2.Determinant)
}

func NewDepedency(determinant set.Set, attributes set.Set) *FunctionalDependency {
	return &FunctionalDependency{determinant, attributes}
}

type Relation struct {
	attributes set.Set
	// We make this a slice and not a set since the order can play a role for some algorithms
	functionalDependencies []*FunctionalDependency
}

func NewRelation(attributes *set.Set, deps []*FunctionalDependency) *Relation {
	return &Relation{
		attributes:             *attributes,
		functionalDependencies: deps,
	}

}

func (r *Relation) Equals(r2 *Relation) bool {
	return r.attributes.Equals(&r2.attributes) && reflect.DeepEqual(r.functionalDependencies, r2.functionalDependencies)
}

func (r *Relation) String() string {
	return getFormattedFunctionalDependencies(r.functionalDependencies)
}

func getFormattedFunctionalDependencies(deps []*FunctionalDependency) string {
	builder := strings.Builder{}
	for _, i := range deps {
		builder.WriteString(fmt.Sprintf("%v -> %v\n", i.Determinant.GetElementsOrdered(), i.Attributes.GetElementsOrdered()))
	}
	return builder.String()
}

// Just a wrapper around Hull
func (r *Relation) Hull(determinants *set.Set) *set.Set {
	return Hull(determinants, r.functionalDependencies)
}

func Hull(determinants *set.Set, functionalDependencies []*FunctionalDependency) *set.Set {
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

func (r *Relation) CandidateKeys() []*set.Set {
	// TODO: check subset of lower keys

	essentialAttributes := r.getEssentialAttributesFromKey()

	foundCandidateKeys := []*set.Set{}

	orderedElements := r.attributes.GetElementsOrdered()

	queue := []*set.Set{essentialAttributes}

	// If the essential attributes define all the attributes ther will be no other candiate keys
	if r.isKey(essentialAttributes) {
		return queue
	}

	for len(queue) != 0 {

		currentBase := queue[0]
		queue = queue[1:]

		index := getHighestIndex(currentBase.GetElementsOrdered(), orderedElements)

		for _, element := range orderedElements[index:] {
			workingCopy := currentBase.DeepCopy()

			workingCopy.Add(element)

			if setIsSubsetOfInSlice(workingCopy, foundCandidateKeys) {
				continue
			}

			if r.isKey(workingCopy) {
				foundCandidateKeys = append(foundCandidateKeys, workingCopy)
			} else {
				queue = append(queue, workingCopy)
			}

		}
	}

	return foundCandidateKeys
}

// returns true if any element of slice is a subset of s
func setIsSubsetOfInSlice(s *set.Set, slice []*set.Set) bool {
	for _, elem := range slice {
		if elem.IsSubSet(s) {
			return true
		}
	}
	return false
}

// expects base to be a sublist of olist2, expects both lists to be ordered
func getHighestIndex(base []string, olist2 []string) int {
	if len(base) == 0 {
		return 0
	}
	return min(slices.Index(olist2, base[len(base)-1])+1, len(olist2))
}

func (r *Relation) isKey(potKey *set.Set) bool {
	return r.Hull(potKey).Equals(&r.attributes)
}

// TODO: make smarter either by better algo or easy fix with dynamic programming

// Returns all the attributes of the relation which are not determined by any other keys
func (r *Relation) getEssentialAttributesFromKey() *set.Set {

	// Atrributes which are only on the left side of rules
	essentialAttributes := r.attributes.DeepCopy()

	for _, rule := range r.functionalDependencies {
		essentialAttributes.Subtract(&rule.Attributes)
	}

	return essentialAttributes
}
