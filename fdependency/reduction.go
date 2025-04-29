package fdependency

import (
	"fmt"
	"slices"
)

func (r *Relation) CanonicalCover() {
	r.leftReduction()
	r.rightReduction()
	r.removeEmptyDependencies()
	r.mergeRulesWithSameDeterminant()
}

func (r *Relation) leftReduction() {

	for _, dependency := range r.functionalDependencies {
		leftReduceDependency(dependency, r)
	}

}

func leftReduceDependency(dependency *FunctionalDependency, r *Relation) {
	for _, attribute := range dependency.Determinant.GetElementsOrdered() {
		newDeterminants := dependency.Determinant.DeepCopy()
		newDeterminants.Remove(attribute)

		newHull := Hull(newDeterminants, r.functionalDependencies)

		if newHull.Equals(Hull(&dependency.Determinant, r.functionalDependencies)) {
			dependency.Determinant = *newDeterminants
		}

	}
}

func (r *Relation) rightReduction() {
	for index, dependency := range r.functionalDependencies {
		rightReduceDependency(dependency, index, r.functionalDependencies)
	}
}

func rightReduceDependency(dependency *FunctionalDependency, index int, functionalDependencies []*FunctionalDependency) {
	sliceCopy := make([]*FunctionalDependency, len(functionalDependencies))
	copy(sliceCopy, functionalDependencies)
	for _, psi := range dependency.Attributes.GetElementsOrdered() {
		newAttributes := dependency.Attributes.DeepCopy()
		newAttributes.Remove(psi)

		sliceCopy[index] = NewDepedency(dependency.Determinant, *newAttributes)

		hull := Hull(&dependency.Determinant, sliceCopy)

		if hull.Contains(psi) {
			dependency.Attributes.Remove(psi)
		}

	}
}

// Removes empty rules from the relation which follow the following format Ψ-> Ø
func (r *Relation) removeEmptyDependencies() {
	// TODO: make more efficient memory allocation wise
	strippedRules := make([]*FunctionalDependency, 0, len(r.functionalDependencies))

	for _, dependency := range r.functionalDependencies {
		if !dependency.Attributes.IsEmpty() {
			strippedRules = append(strippedRules, dependency)
		}
	}

	r.functionalDependencies = strippedRules

}

func (r *Relation) mergeRulesWithSameDeterminant() {
	// TODO: make more efficient memory allocation wise
	mergedRules := make([]*FunctionalDependency, 0, len(r.functionalDependencies))

	// TODO: make more efficient time complexity wise
OUTER:
	for _, rule := range r.functionalDependencies {
		for _, mergedRule := range mergedRules {
			if mergedRule.Determinant.Equals(&rule.Determinant) {
				mergedRule.Attributes.AddUnion(&rule.Attributes)
				continue OUTER
			}
		}

		mergedRules = append(mergedRules, rule)

	}

	r.functionalDependencies = mergedRules

}

func (r *Relation) addKeysToFunctionalDependency() {
	for _, v := range r.functionalDependencies {
		v.Attributes.AddUnion(&v.Determinant)
	}
}

// Removes functional dependencies whose total attributes (det + att) are the subset of another fd
func (r *Relation) removeFDwithSameSet() {

	i := 0

	// This code is very inefficient because premature optimization is the root of all evil
OUTER:
	for i < len(r.functionalDependencies) {
		fd := r.functionalDependencies[i]

		for _, fd2 := range r.functionalDependencies {
			if fd2 == fd {
				continue
			}
			if fd.getFullSet().IsSubSet(fd2.getFullSet()) {
				fmt.Println("Removing ", i)
				r.functionalDependencies = slices.Delete(r.functionalDependencies, i, i+1)
				continue OUTER
			}
		}
		i++
	}

}
func removeFromSliceInefficient(slice []*FunctionalDependency, s int) []*FunctionalDependency {
	return append(slice[:s], slice[s+1:]...)
}
