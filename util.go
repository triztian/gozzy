package gozzy

// Generate a linear space within the given range
func linspace(a, b, n int) []float64 {
	var s []float64

	a0 := a
	if a0 == 0 {
		a0 = 1
	}

	for i := a0; i < b; i++ {
		s[i] = 1.0 / float64(i)
	}

	return s
}

// Merge two functions
func funcMerge(fa, fb Mf, bin func(float64, float64) float64) Mf {
	var m Mf = func(x float64) float64 {
		return bin(fa(x), fb(x))
	}

	return m
}
