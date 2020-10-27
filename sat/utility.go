package sat

// Utility

// Multiply returns the cartesian product of the given Literal slices
func Multiply(sets ...Literals) []Literals {
	if len(sets) == 0 {
		return sets
	}

	var currentProduct []Literals = make([]Literals, 0)
	// sets[0] = {a_1, a_2, ...}
	// currentProduct = {{a_1}, {a_2}, ...}
	for _, literal := range sets[0] {
		currentProduct = append(currentProduct, []Literal{literal})
	}

	// A = currentProduct = {{a_aindex * b_bindex * ... n_nindex}}
	// B = {m_1, m_2, ...}
	for _, set := range sets[1:] {
		nextProduct := make([]Literals, 0)
		for _, a := range currentProduct { // type(a) = Literals = []Literal
			for _, b := range set { // type(b) = Literal
				nextProduct = append(nextProduct, append(a, b))
			}
		}
		currentProduct = nextProduct
	}

	return currentProduct
}

// ExactlyOneTrue returns clauses specifying that exactly one of the given literals is true.
func ExactlyOneTrue(literals []Literal) ConjunctiveFormula {
	clauses := make([]DisjunctiveClause, 0)
	// At least one literal is true.
	clauses = append(clauses, NewDisjunctiveClause(literals...))

	for i, litA := range literals {
		for _, litB := range literals[i+1:] {
			// No two literals are true.
			clauses = append(clauses, NewDisjunctiveClause(litA.Negate(), litB.Negate()))
		}
	}

	return NewConjunctiveFormula(clauses)
}
