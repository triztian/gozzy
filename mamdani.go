package gozzy

type MamdaniT1 struct {
	name     string
	inputs   map[string]Range
	outputs  map[string]Range
	terms    map[string]FuzzySetT1
	rules    []Rule
	defuzzer Defuzzer
}

// Process the given perceptions
func (s MamdaniT1) process(values Crisp) Crisp {
	var sets []FuzzySetT1

	for _, rule := range s.rules {
		sets = append(sets, s.processRule(rule, values))
	}

	return s.defuzz(sets)
}

// Process a system rule with the given perceptions
func (s MamdaniT1) processRule(rule Rule, values Crisp) FuzzySetT1 {
	res := rule.antecedent.evaluate(values)
	return res
}

func (s MamdaniT1) defuzz(sets []FuzzySetT1) Crisp {
	var response Crisp

	return response
}

// Obtain the alpha cuts of the fuzzy set.
// An alpha cut is defined as the membership value that
// is the same for two **adjecent** elements.
func (z FuzzySetT1) alphas() []float64 {
	var a []float64
	for i, x := range z.U[1:] {
		mfx := z.m(z.U[i-1])
		if mfx == z.m(x) {
			a = append(a, mfx)
		}
	}
	return a
}
