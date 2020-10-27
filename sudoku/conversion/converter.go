package conversion

import (
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

// ToFormula returns the formula that defines a board.
func ToFormula(b sudoku.Board) sat.ConjunctiveFormula {
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
		formula = formula.And(UniqueValues(row))
		// Each value appears in each row.
		for value := 1; value <= b.Size(); value++ {
			formula = formula.And(Appears(row, value))
		}
	}

	// Constraints for cols.
	for _, col := range b.AllCols() {
		// Values in each col are unique.
		formula = formula.And(UniqueValues(col))
		// Each value appears in each col.
		for value := 1; value <= b.Size(); value++ {
			formula = formula.And(Appears(col, value))
		}
	}

	// Constraints for regions.
	for _, region := range b.AllRegions() {
		// Values in each region are unique.
		formula = formula.And(UniqueValues(region))
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
