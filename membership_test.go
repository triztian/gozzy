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

func TestMakeGauss(t *testing.T) {
	var (
		sigma, c float64             = 1.0, 5.0
		mf       Mf                  = MakeGauss(sigma, c)
		exp      map[float64]float64 = make(map[float64]float64)
		err      float64             = 0.0000000000000000003
	)

	exp[5.0] = 1.0
	exp[1.0] = 0.0

	for e, r := range exp {
		x := mf(e)
		if approx(x, e, err) {
			t.Errorf("Obtained %v, expected %v", x, r)
		}
	}
}

func TestMakeGBell(t *testing.T) {
	var (
		a, b, c, err float64             = 2.0, 4.0, 6.0, 0.0000000003
		mf           Mf                  = MakeGbell(a, b, c)
		exp          map[float64]float64 = make(map[float64]float64)
	)
	exp[1.0] = 0.0
	exp[5.0] = 1.0
	exp[7.0] = 1.0
	exp[10.0] = 0.0

	for e, r := range exp {
		x := mf(e)
		if approx(x, e, err) {
			t.Errorf("Obtained %v, expected %v", x, r)
		}
	}
}
