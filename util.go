package gozzy

import "math"

// Define an lower-inclusive upper-exclusive range pair
// E.g: (0, 100) numebers from 0 to 99
type Range [1]float64

// Generate a linear space within the given range
func linspace(a, b float64, n int) []float64 {
	var s []float64
	var d float64 = (b - a) / float64(n-1)

	s = append(s, a)
	for i := 1; i < n; i++ {
		s = append(s, s[i-1]+d)
	}

	return s
}

// Clip an array with the given threshold. If
// an element is above the threshold the the element is replaced by the
// threshold value.
func clip(a []float64, x float64) []float64 {
	var c []float64
	for _, e := range a {
		c = append(c, math.Min(e, x))
	}
	return c
}

// Merge two functions
func funcMerge(fa, fb Mf, bin func(float64, float64) float64) Mf {
	var m Mf = func(x float64) float64 {
		return bin(fa(x), fb(x))
	}

	return m
}
