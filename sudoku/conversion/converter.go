package conversion

import "fmt"

// ParseState parses boolean state back into a board.
func ParseState(state map[string]bool) Board {
	initialValues := make(map[Coordinate]int)
	for name, v := range state {
		if v {
			coordinate, value := ParseName(name)
			initialValues[coordinate] = value
		}
	}
	return NewStandardBoard(initialValues)
}

// Diagonal returns the coordinates in the specified diagonal, in order.
func (b Board) Diagonal(start, end Coordinate) (coordinates Coordinates) {
	rowSign := 1
	if end.row < start.row {
		rowSign = -1
	}
	colSign := 1
	if end.col < start.col {
		colSign = -1
	}

	rows := (end.row - start.row) * rowSign
	cols := (end.col - start.col) * colSign
	if rows != cols {
		panic(fmt.Sprintf("Not on a diagonal! %s, %s", start, end))
	}

	for i := 0; i <= rows; i++ {
		row := start.row + (rowSign * i)
		col := start.col + (colSign * i)
		coordinates = append(coordinates, NewCoordinate(row, col))
	}

	return coordinates
}

// Formula returns the formula defining the board.
func (b Board) Formula() ConjunctiveFormula {
	formula := EmptyConjunctiveFormula()

	// Constraints for individual cells.
	clauses := make([]DisjunctiveClause, 0)
	for _, coordinate := range b.GetAllCoordinates() {
		clauses = append(clauses, coordinate.Clauses()...)

		if value, ok := b.values[coordinate]; ok {
			clause := NewDisjunctiveClause(coordinate.Literal(value))
			clauses = append(clauses, clause)
		}
	}
	formula = formula.And(ConjunctiveFormula{clauses})

	// Constraints for rows, cols, and regions.
	clauses = make([]DisjunctiveClause, 0)
	for i := 1; i <= 9; i++ {
		row := b.Row(i)
		col := b.Col(i)
		region := b.Region((i-1)/3+1, (i-1)%3+1)

		formula = formula.And(UniqueValues(row)).And(Appears(row, allValues...))
		formula = formula.And(UniqueValues(col)).And(Appears(col, allValues...))
		formula = formula.And(UniqueValues(region)).And(Appears(region, allValues...))
	}

	// Additional rules.
	if b.rules.diagonalsUnique {
		diagonal := b.Diagonal(NewCoordinate(1, 1), NewCoordinate(9, 9))
		formula = formula.And(UniqueValues(diagonal)).And(Appears(diagonal, allValues...))

		diagonal = b.Diagonal(NewCoordinate(1, 9), NewCoordinate(9, 1))
		formula = formula.And(UniqueValues(diagonal)).And(Appears(diagonal, allValues...))
	}

	return formula
}
