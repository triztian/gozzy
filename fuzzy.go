package gozzy

import "math"

type Crisp map[string]float64
type Operator func(a, b FuzzySetT1) FuzzySetT1

type Exp struct {
	p  *FuzzySetT1
	op *Operation
	e  *Exp
}

type Operation struct {
	operator Operator
	a, b     FuzzySetT1
}

// A Structure representing a Fuzzy Set Type-I
// As defined by wikipedia:
// https://en.wikipedia.org/wiki/Fuzzy_set
type FuzzySetT1 struct {
	term string
	U    []float64
	m    Mf
}

func (op Operation) resolve() FuzzySetT1 {
	return op.operator(op.a, op.b)
}

// Evaluate the expression with the given context
func (e Exp) evaluate(context Crisp) FuzzySetT1 {

	if e.e == nil && e.op == nil && e.p != nil {
		if val, ok := context[e.p.term]; ok {
			return Clip(*e.p, e.p.m(val))
		} else {
			panic("in the disco")
		}
	} else if e.e == nil && e.p == nil && e.op != nil {
		return e.op.resolve()

	} else if e.op == nil && e.p == nil && e.e != nil {
		return e.e.evaluate(context)
	} else {
		panic("Incorrect expression form")
	}
}

func Union(a, b FuzzySetT1) FuzzySetT1 {
	m := funcMerge(a.m, b.m, math.Max)
	return FuzzySetT1{a.term + b.term, a.U, m}
}

func Intersection(a, b FuzzySetT1) FuzzySetT1 {
	m := funcMerge(a.m, b.m, math.Min)

	return FuzzySetT1{a.term + b.term, a.U, m}
}

// Obtain the complement of a Fuzzy Set Type-I
func Complement(a FuzzySetT1) FuzzySetT1 {
	not_mf := func(x float64) float64 {
		return 1.0 - a.m(x)
	}

	return FuzzySetT1{a.term + "~", a.U, not_mf}
}

// Clip a fuzzy set to the given bound
func Clip(z FuzzySetT1, x float64) FuzzySetT1 {
	var elems []float64
	for _, e := range z.U {
		if z.m(e) <= x {
			elems = append(elems, e)
		}
	}
	return FuzzySetT1{z.term, elems, z.m}
}

// Obtain the memberships of a FuzzySet
func (z FuzzySetT1) memberships() []float64 {
	var mems []float64
	for _, x := range z.U {
		mems = append(mems, z.m(x))
	}
	return mems
}

// Obtain the supreme element and it's membership value
func (z FuzzySetT1) supreme() (float64, float64) {
	var sup_x float64
	sup_m := z.m(z.U[0])

	for _, x := range z.U[1:len(z.U)] {
		mfx := z.m(x)
		if mfx > sup_m {
			sup_x = x
			sup_m = mfx
		}
	}

	return sup_x, sup_m
}

// Obtain the alpha cuts of the fuzzy set.
// An alpha cut is defined as the membership value that
// is the same for two **adjecent** elements.
func (z FuzzySetT1) alphas() []float64 {
	var a []float64
	for i, x := range z.U[1:] {
		mfx := z.m(z.U[i-1])
		if mfx == z.m(x) {
			a = append(a, mfx)
		}
	}
	return a
}

// A fuzzy system rule. Rules have the form of:
// "IF temp IS hot AND energy IS high THEN fan IS on, radiator IS empty"
// At the moment rules are limited to provide the values for a single perception.
// which in our example is the perception "fan"
type Rule struct {
	antecedent Exp
	consecuent map[string]Exp
}

type MamdamiT1 struct {
	name     string
	inputs   map[string]Range
	outputs  map[string]Range
	terms    map[string]FuzzySetT1
	rules    []Rule
	defuzzer Defuzzer
}

// Process the given perceptions
func (s MamdamiT1) process(values Crisp) Crisp {
	var sets []FuzzySetT1

	for _, rule := range s.rules {
		sets = append(sets, s.processRule(rule, values))
	}

	return s.defuzz(sets)
}

// Process a system rule with the given perceptions
func (s MamdamiT1) processRule(rule Rule, values Crisp) FuzzySetT1 {
	res := rule.antecedent.evaluate(values)
	return res
}

func (s MamdamiT1) defuzz(sets []FuzzySetT1) Crisp {
	var response Crisp

	return response
}
