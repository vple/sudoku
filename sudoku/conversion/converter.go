package conversion

import (
	"fmt"

	sudoku ".."
	"../../sat"
)

// ParseState parses boolean state back into a board.
func ParseState(state map[string]bool) sudoku.Board {
	initialValues := make(map[sudoku.Coordinate]int)
	for name, v := range state {
		if v {
			coordinate, value := fromName(name)
			initialValues[coordinate] = value
		}
	}
	return sudoku.NewStandardBoard(initialValues)
}

// ToFormula converts a board to CNF form.
func ToFormula(board sudoku.Board) (formula sat.ConjunctiveFormula) {
	constraints := board.AllConstraints()

	for _, constraint := range constraints {
		formula = formula.And(convert(constraint, board))
	}

	return formula
}

func convert(c sudoku.Constraint, board sudoku.Board) sat.ConjunctiveFormula {
	switch constraint := c.(type) {
	case sudoku.CellValueConstraint:
		return constrainToValues(constraint.Coordinate(), constraint.Values())
	case sudoku.UniqueValueConstraint:
		return uniqueValues(constraint.Coordinates(), board.AllValues())
	case sudoku.ContainsValuesConstraint:
		return Appears(constraint.Coordinates(), constraint.Values()...)
	case sudoku.IncreasingValueConstraint:
		formula := sat.EmptyConjunctiveFormula()

		clauses := make([]sat.DisjunctiveClause, 0)
		coordinates := constraint.Coordinates()
		size := board.Size()
		dof := size - len(coordinates) // Degrees of freedom
		for i, a := range coordinates {
			possibleAValues := make([]int, 0)
			for x := 0; x <= dof; x++ {
				possibleAValues = append(possibleAValues, i+x+1)
			}
			formula = formula.And(constrainToValues(a, possibleAValues))

			for _, b := range coordinates[i+1:] {
				// Coordinate A < Coordinate B
				for _, aValue := range board.AllValues() {
					for bValue := 1; bValue < aValue; bValue++ {
						// For aValue > bValue, we want:
						// !aValue || !bValue
						notA := toLiteral(a, aValue).Negate()
						notB := toLiteral(b, bValue).Negate()
						clauses = append(clauses, sat.NewDisjunctiveClause(notA, notB))
					}
				}
			}
		}

		return formula.And(sat.NewConjunctiveFormula(clauses))
	default:
		panic(fmt.Sprintf("Unknown constraint type: %T", c))
	}
}

func constrainToValues(coordinate sudoku.Coordinate, values []int) sat.ConjunctiveFormula {
	literals := make([]sat.Literal, 0)
	for _, value := range values {
		literals = append(literals, toLiteral(coordinate, value))
	}
	return sat.ExactlyOneTrue(literals)
}
