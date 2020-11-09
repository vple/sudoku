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
		literals := make([]sat.Literal, 0)
		for _, value := range constraint.Values() {
			literals = append(literals, toLiteral(constraint.Coordinate(), value))
		}
		return sat.ExactlyOneTrue(literals)
	case sudoku.UniqueValueConstraint:
		return uniqueValues(constraint.Coordinates(), board.AllValues())
	case sudoku.ContainsValuesConstraint:
		return Appears(constraint.Coordinates(), constraint.Values()...)
	default:
		panic(fmt.Sprintf("Unknown constraint type: %T", c))
	}
}
