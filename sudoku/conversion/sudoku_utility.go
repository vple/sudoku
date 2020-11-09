package conversion

import (
	sudoku ".."
	"../../sat"
)

// AppearsOnce specifies that the given value appears at exactly one of the given coordinates.
func AppearsOnce(coordinates []sudoku.Coordinate, value int) sat.ConjunctiveFormula {
	literals := make([]sat.Literal, 0)
	for _, coordinate := range coordinates {
		literal := toLiteral(coordinate, value)
		literals = append(literals, literal)
	}

	return sat.ExactlyOneTrue(literals)
}

// Appears specifies that each of the values appears at least one of the given coordinates.
func Appears(coordinates []sudoku.Coordinate, values ...int) sat.ConjunctiveFormula {
	clauses := make([]sat.DisjunctiveClause, 0)
	for _, value := range values {
		literals := make(sat.Literals, 0)
		for _, coordinate := range coordinates {
			literal := toLiteral(coordinate, value)
			literals = append(literals, literal)
		}
		clauses = append(clauses, sat.NewDisjunctiveClause(literals...))
	}
	return sat.NewConjunctiveFormula(clauses)
}

// uniqueValues specifies that all coordinates have unique values.
func uniqueValues(coordinates []sudoku.Coordinate, possibleValues []int) sat.ConjunctiveFormula {
	clauses := make([]sat.DisjunctiveClause, 0)
	for i, a := range coordinates {
		for _, b := range coordinates[i+1:] {
			for _, value := range possibleValues {
				// a != value || b != value
				notA := toLiteral(a, value).Negate()
				notB := toLiteral(b, value).Negate()
				clause := sat.NewDisjunctiveClause(notA, notB)
				clauses = append(clauses, clause)
			}
		}
	}
	return sat.NewConjunctiveFormula(clauses)
}

// Sum sums the values of these literals as if they were true.
func sum(literals sat.Literals) (sum int) {
	for _, literal := range literals {
		l := literal.(sat.PositiveLiteral)
		_, value := fromName(l.Name())
		sum += value
	}
	return sum
}

// SumEquals returns the formula specifying that the sums of the values of the given coordinates equals the desired sum.
func SumEquals(coordinates sudoku.Coordinates, total int) sat.ConjunctiveFormula {
	var allCombinations []sat.Literals = allPermutations(coordinates)
	clauses := make([]sat.DisjunctiveClause, 0)
	for _, combination := range allCombinations {
		if sum(combination) == total {
			clauses = append(clauses, sat.NewDisjunctiveClause(combination...))
		}
	}

	return sat.NewConjunctiveFormula(clauses)
}

// allPermutations determines all possible permuations for the specified cells.
func allPermutations(coordinates sudoku.Coordinates) []sat.Literals {
	literals := make([]sat.Literals, 0)
	for _, coordinate := range coordinates {
		literals = append(literals, toLiterals(coordinate))
	}
	return sat.Multiply(literals...)
}
