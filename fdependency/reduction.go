package fdependency

import ()

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

func leftReduceDependency(dependency *functionalDependency, r *Relation) {
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

func rightReduceDependency(dependency *functionalDependency, index int, functionalDependencies []*functionalDependency) {
	sliceCopy := make([]*functionalDependency, len(functionalDependencies))
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
	strippedRules := make([]*functionalDependency, 0, len(r.functionalDependencies))

	for _, dependency := range r.functionalDependencies {
		if !dependency.Attributes.IsEmpty() {
			strippedRules = append(strippedRules, dependency)
		}
	}

	r.functionalDependencies = strippedRules

}

func (r *Relation) mergeRulesWithSameDeterminant() {
	// TODO: make more efficient memory allocation wise
	mergedRules := make([]*functionalDependency, 0, len(r.functionalDependencies))

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
