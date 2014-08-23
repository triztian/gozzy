package gozzy

import "math"

// A Structure representing a Fuzzy Set Type-I
type FuzzySetT1 struct {
	elems []float32
	mf    Mf
}

func Union(a, b FuzzySetT1) FuzzySetT1 {
	m := funcMerge(a.mf, b.mf, math.Max)

	return FuzzySetT1{a.elems, m}
}

func Intersection(a, b FuzzySetT1) FuzzySetT1 {
	m := funcMerge(a.mf, b.mf, math.Min)

	return FuzzySetT1{a.elems, m}
}

// Obtain the complement of a Fuzzy Set Type-I
func Complement(a FuzzySetT1) FuzzySetT1 {
	not_mf := func(x float64) float64 {
		return 1.0 - a.mf(x)
	}

	return FuzzySetT1{a.elems, not_mf}
}
