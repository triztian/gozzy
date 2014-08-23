package gozzy

type Defuzzer func(sets []FuzzySetT1) float64

// The centroid defuzzer
func Centroid(sets []FuzzySetT1) float64 {
	var (
		area, t float64 = 0.0, 0.0
		y       FuzzySetT1
	)

	y = sets[0]
	for _, z := range sets[1:len(sets)] {
		y = Union(y, z)
	}

	for _, x := range y.elems {
		mfx := y.mf(x)
		area += mfx
		t += mfx * x
	}
	return t / area
}

// The height defuzzer
func Height(sets []FuzzySetT1) float64 {
	var (
		memberships map[float64]float64
		s           float64
		singleton   FuzzySetT1
	)

	for _, z := range sets {
		s, memberships[s] = z.supreme()
	}

	singleton = FuzzySetT1{"", sets[0].elems, MakeMapped(memberships)}
	return Centroid([]FuzzySetT1{})
}
