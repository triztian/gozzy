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

	for _, x := range y.U {
		mfx := y.m(x)
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
		symbol      string
	)

	for _, z := range sets {
		s, memberships[s] = z.supreme()
		symbol += z.term
	}

	singleton = FuzzySetT1{symbol, sets[0].U, MakeMapped(memberships)}
	return Centroid([]FuzzySetT1{singleton})
}
