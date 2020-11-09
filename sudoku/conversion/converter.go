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

func ToFormula(board sudoku.Board) (formula sat.ConjunctiveFormula) {
	constraints := board.AllConstraints()

	for _, constraint := range constraints {
		// fmt.Printf("%T: %v\n", constraint, constraint)
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

// ToFormula returns the formula that defines a board.
func ToFormula2(b sudoku.Board) sat.ConjunctiveFormula {
	formula := sat.EmptyConjunctiveFormula()

	// Constraints for individual cells.
	// clauses := make([]sat.DisjunctiveClause, 0)
	for _, coordinate := range b.AllCoordinates() {
		// Coordinate has exactly one value from 1-9.
		hasOneValue := sat.ExactlyOneTrue(toLiterals(coordinate))
		formula = formula.And(hasOneValue)

		// clauses = append(clauses, coordinate.Clauses()...)

		if value, ok := b.Value(coordinate); ok {
			// If value is known, specify it.
			exactValue := sat.NewDisjunctiveClause(toLiteral(coordinate, value))
			formula = formula.And(exactValue.ToFormula())
			// clauses = append(clauses, clause)
		}
	}
	// formula = formula.And(sat.ConjunctiveFormula{clauses})

	// Constraints for rows.
	for _, row := range b.AllRows() {
		// Values in each row are unique.
		formula = formula.And(uniqueValues(row, b.AllValues()))
		// Each value appears in each row.
		for value := 1; value <= b.Size(); value++ {
			formula = formula.And(Appears(row, value))
		}
	}

	// Constraints for cols.
	for _, col := range b.AllCols() {
		// Values in each col are unique.
		formula = formula.And(uniqueValues(col, b.AllValues()))
		// Each value appears in each col.
		for value := 1; value <= b.Size(); value++ {
			formula = formula.And(Appears(col, value))
		}
	}

	// Constraints for regions.
	for _, region := range b.AllRegions() {
		// Values in each region are unique.
		formula = formula.And(uniqueValues(region, b.AllValues()))
		// Each value appears in each region.
		for value := 1; value <= b.Size(); value++ {
			formula = formula.And(Appears(region, value))
		}
	}

	// for i := 1; i <= b.Size(); i++ {
	// 	row := b.Row(i)
	// 	col := b.Col(i)
	// 	// Currently assumes size is 9.
	// 	region := b.Region((i-1)/3+1, (i-1)%3+1)

	// 	formula = formula.And(UniqueValues(row)).And(Appears(row, allValues...))
	// 	formula = formula.And(UniqueValues(col)).And(Appears(col, allValues...))
	// 	formula = formula.And(UniqueValues(region)).And(Appears(region, allValues...))
	// }

	// Additional rules.
	// if b.rules.diagonalsUnique {
	// 	diagonal := b.Diagonal(sudoku.NewCoordinate(1, 1), sudoku.NewCoordinate(9, 9))
	// 	formula = formula.And(UniqueValues(diagonal)).And(Appears(diagonal, allValues...))

	// 	diagonal = b.Diagonal(NewCoordinate(1, 9), NewCoordinate(9, 1))
	// 	formula = formula.And(UniqueValues(diagonal)).And(Appears(diagonal, allValues...))
	// }

	return formula
}
