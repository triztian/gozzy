package gozzy

import "math"

type Crisp map[string]float64
type FuzzyOperator func(a, b FuzzySetT1) FuzzySetT1

type FuzzyExpression interface {
	resolve(values Crisp) FuzzySetT1
}

// A Structure representing a Fuzzy Set Type-I
type FuzzySetT1 struct {
	perceptor string
	elems     []float64
	mf        Mf
}

func (z FuzzySetT1) resolve(values Crisp) FuzzySetT1 {
	var r FuzzySetT1
	for p, v := range values {
		if z.perceptor == p {
			r = Clip(z, z.mf(v))
			return r
		}
	}

	return z
}

func Union(a, b FuzzySetT1) FuzzySetT1 {
	m := funcMerge(a.mf, b.mf, math.Max)

	return FuzzySetT1{a.perceptor, a.elems, m}
}

func Intersection(a, b FuzzySetT1) FuzzySetT1 {
	m := funcMerge(a.mf, b.mf, math.Min)

	return FuzzySetT1{a.perceptor, a.elems, m}
}

// Obtain the complement of a Fuzzy Set Type-I
func Complement(a FuzzySetT1) FuzzySetT1 {
	not_mf := func(x float64) float64 {
		return 1.0 - a.mf(x)
	}

	return FuzzySetT1{a.perceptor, a.elems, not_mf}
}

// Clip a fuzzy set to the given bound
func Clip(z FuzzySetT1, x float64) FuzzySetT1 {
	var (
		elems []float64
	)
	for _, e := range z.elems {
		if z.mf(e) <= x {
			elems = append(elems, e)
		}
	}
	return FuzzySetT1{z.perceptor, elems, z.mf}
}

// Obtain the memberships of a FuzzySet
func (z FuzzySetT1) memberships() []float64 {
	var mems []float64
	for i, x := range z.elems {
		mems = append(mems, z.mf(x))
	}
	return mems
}

// Obtain the supreme element and it's membership value
func (z FuzzySetT1) supreme() (float64, float64) {
	var sup_x float64
	sup_m := z.mf(z.elems[0])

	for _, x := range z.elems[1:len(z.elems)] {
		mfx := z.mf(x)
		if mfx > sup_m {
			sup_x = x
			sup_m = mfx
		}
	}

	return sup_x, sup_m
}

// Obtaint the alpha cuts of the fuzzy set.
// An alpha cut is defined as the membership value that
// is the same for two **adjecent** elements.
func (z FuzzySetT1) alphas() []float64 {
	var a []float64
	for i, x := range z.elems[1:] {
		mfx := z.mf(z.elems[i-1])
		if mfx == z.mf(z.elems[i]) {
			a = append(a, mfx)
		}
	}
	return a
}

type FuzzyOperation struct {
	operator FuzzyOperator
	a, b     FuzzyExpression
}

func (e FuzzyOperation) resolve(values Crisp) FuzzySetT1 {
	return e.operator(e.a.resolve(values), e.b.resolve(values))
}

type ExpressionBuilder struct {
	_expression FuzzyExpression
}

func (eb ExpressionBuilder) addExpression(op FuzzyOperator, terms ...FuzzyExpression) ExpressionBuilder {
	for _, term := range terms {
		eb._expression = FuzzyOperation{op, eb._expression, term}
	}
	return eb
}

func (eb ExpressionBuilder) and(terms ...FuzzyExpression) ExpressionBuilder {
	return eb.addExpression(Intersection, terms...)
}

func (eb ExpressionBuilder) Or(terms ...FuzzyExpression) ExpressionBuilder {
	return eb.addExpression(Union, terms...)
}

type Rule struct {
	antecedent, consecuent FuzzyExpression
}

type MamdamiT1 struct {
	name     string
	inputs   map[string]Range
	outputs  map[string]Range
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
	ant := rule.antecedent.resolve(values)
	alphas := ant.alphas()
	cons := rule.consecuent.resolve(values)
	return z
}

func (s MamdamiT1) defuzz(sets []FuzzySetT1) Crisp {
	return s.defuzzer(sets)
}
