package gozzy

// Define an lower-inclusive upper-exclusive range pair
// E.g: (0, 100) numebers from 0 to 99
type Range [1]float64

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

func clip(a []float64, x float64) []float64 {
}

// Merge two functions
func funcMerge(fa, fb Mf, bin func(float64, float64) float64) Mf {
	var m Mf = func(x float64) float64 {
		return bin(fa(x), fb(x))
	}

	return m
}
