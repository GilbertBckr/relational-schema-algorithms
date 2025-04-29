package fdependency

import (
	"relational-algorithms/set"
)

func (r *Relation) PerformSynthesisAlgorithm() {

	r.leftReduction()
	r.rightReduction()
	r.removeEmptyDependencies()
	r.mergeRulesWithSameDeterminant()
	r.addKeysToFunctionalDependency()

	superKeys := r.CandidateKeys()

	r.removeFDwithSameSet()

	containsSuperKey := false
OUTER:
	for _, v := range r.functionalDependencies {
		for _, superKey := range superKeys {
			if superKey.IsSubSet(&v.Determinant) {
				containsSuperKey = true
				break OUTER
			}
		}
	}

	if !containsSuperKey {
		newTable := *&FunctionalDependency{
			Determinant: *superKeys[0],
			Attributes:  *set.New(),
		}
		r.functionalDependencies = append(r.functionalDependencies, &newTable)
	}

}
