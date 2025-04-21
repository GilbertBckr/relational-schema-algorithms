package fdependency

func (r *Relation) LeftReduction() {

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
