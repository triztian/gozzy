package gozzy

import "math"

type Mf func(float64) float64

// The Gauss membership function maker
//
func MakeGauss(sigma, c float64) Mf {
	return func(x float64) float64 {
		return float64(math.Exp(-(math.Pow(x-c, 2.0)) / (2.0 * math.Pow(sigma, 2.0))))
	}
}

// Maker function of the GBell membership function
func MakeGbell(a, b, c float64) Mf {
	return func(x float64) float64 {
		return 1.0 / (1.0 + math.Pow(math.Abs((x-c)/a), 2*b))
	}
}

// Maker funciton of the trapezoid membership function
func MakeTrapezoid(a, b, c, d float64) Mf {
	return func(x float64) float64 {
		return math.Max(math.Min(math.Min(((x-a)/(b-a)), 1), (d-x)/(d-c)), 0.0)
	}
}

// Maker funciton of the Triangular membership function
func MakeTriangular(a, b, c float64) Mf {
	return func(x float64) float64 {
		return math.Max(math.Min(((x-a)/(b-a)), ((c-x)/(c-b))), 0)
	}
}

// Maker function of the Mapped function, the mapped function
// is used when there is not a way to know how the membership of an item
// was calculated but the membership is known.
func MakeMapped(m map[float64]float64) Mf {
	return func(x float64) float64 {
		return m[x]
	}
}
