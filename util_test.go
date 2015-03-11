package gozzy

import (
	"math"
	"testing"
)

func TestLinspace(t *testing.T) {
	a, b := 0.0, 5.0
	n := 10
	expected := []float64{0.0, 0.5, 1.0, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5.0}
	lin := linspace(a, b, n)

	if len(lin) != len(expected) {
		t.Error("Unexpected length: %d", len(lin))
	}

	for i := 0; i < len(lin); i++ {
		if lin[i] != expected[i] {
			t.Error("Unexpected element: %v -> %v, should be %v", i, lin[i], expected[i])
		}
	}
}

func TestClip(t *testing.T) {
	x := 5.0
	a := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	e := []float64{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 3, 2, 1}
	r := clip(a, x)

	if len(e) != len(e) {
		t.Error("Unexpected length: %d", len(r))
	}

	for i := 0; i < len(r); i++ {
		if r[i] != e[i] {
			t.Error("Unexpected element: %v -> %v, should be %v", i, r[i], e[i])
		}
	}
}

func TestFuncMerge(t *testing.T) {
	x := 1.0
	bin := math.Max
	fa := func(x float64) float64 {
		return x + 1
	}

	fb := func(x float64) float64 {
		return x + 100
	}
	e := 101.0

	fr := funcMerge(fa, fb, bin)
	r := fr(x)

	if r != e {
		t.Error("Unexpected result: %v should be %v", r, e)
	}
}
