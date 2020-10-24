package conversion

import "../../sat"

// AppearsOnce specifies that the given value appears at exactly one of the given coordinates.
func AppearsOnce(coordinates []Coordinate, value int) []sat.DisjunctiveClause {
	literals := make([]sat.Literal, 0)
	for _, coordinate := range coordinates {
		literal := coordinate.Literal(value)
		literals = append(literals, literal)
	}

	return sat.ExactlyOneTrue(literals)
}

// Appears specifies that each of the values appears at least one of the given coordinates.
func Appears(coordinates []Coordinate, values ...int) sat.ConjunctiveFormula {
	clauses := make([]sat.DisjunctiveClause, 0)
	for _, value := range values {
		literals := make(sat.Literals, 0)
		for _, coordinate := range coordinates {
			literal := coordinate.Literal(value)
			literals = append(literals, literal)
		}
		clauses = append(clauses, sat.NewDisjunctiveClause(literals...))
	}
	return sat.NewConjunctiveFormula(clauses)
}

// UniqueValues specifies that all coordinates have unique values.
func UniqueValues(coordinates []Coordinate) sat.ConjunctiveFormula {
	clauses := make([]sat.DisjunctiveClause, 0)
	for i, a := range coordinates {
		for _, b := range coordinates[i+1:] {
			for value := 1; value <= 9; value++ {
				// a != value || b != value
				clause := NewDisjunctiveClause(a.Literal(value).Negate(), b.Literal(value).Negate())
				clauses = append(clauses, clause)
			}
		}
	}
	return ConjunctiveFormula{clauses}
}

// Sum sums the values of these literals as if they were true.
func sum(literals sat.Literals) (sum int) {
	for _, literal := range literals {
		l := literal.(sat.PositiveLiteral)
		sum += l.GetValue()
	}
	return sum
}

// SumEquals returns the formula specifying that the sums of the values of the given coordinates equals the desired sum.
func SumEquals(coordinates Coordinates, sum int) sat.ConjunctiveFormula {
	var allCombinations []Literals = allPermutations(coordinates)
	clauses := make([]DisjunctiveClause, 0)
	for _, combination := range allCombinations {
		if combination.Sum() == sum {
			clauses = append(clauses, DisjunctiveClause{combination})
		}
	}

	return ConjunctiveFormula{clauses}
}

// allPermutations determines all possible permuations for the specified cells.
func allPermutations(coordinates Coordinates) []sat.Literals {
	literals := make([]Literals, 0)
	for _, coordinate := range coordinates {
		literals = append(literals, coordinate.Literals())
	}
	return Multiply(literals...)
}
