package gozzy

import "testing"

func TestLinspace(t *testing.T) {
	a, b := 1.0, 10.0
	n := 10
	lin := linspace(a, b, n)

	t.Log("Linspace (%f, %f, %v): %v\n", a, b, n, lin)
	t.Fail()
}
