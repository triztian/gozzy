package gozzy

import "testing"

func TestMakeTriangular(t *testing.T) {
	var (
		a, b, c float64             = 0.0, 0.5, 1.0
		mf      Mf                  = MakeTriangular(a, b, c)
		exp     map[float64]float64 = make(map[float64]float64)
	)

	exp[0.0] = 0.0
	exp[0.5] = 1.0
	exp[1.0] = 0.0

	for e, r := range exp {
		x := mf(e)
		if x != r {
			t.Errorf("Obtained %v, expected %v", x, r)
		}
	}
}

func TestMakeTrapezoid(t *testing.T) {
	var (
		a, b, c, d float64             = 0.0, 0.25, 0.75, 1.0
		mf                             = MakeTrapezoid(a, b, c, d)
		exp        map[float64]float64 = make(map[float64]float64)
	)
	exp[0.0] = 0.0
	exp[0.25] = 1.0
	exp[0.75] = 1.0
	exp[1.0] = 0.0

	for e, r := range exp {
		x := mf(e)
		if x != r {
			t.Errorf("Obtained %v, expected %v", x, r)
		}
	}
}
