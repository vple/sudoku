package sudoku

import (
	"fmt"
	"strconv"
	"strings"
)

var allValues = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Board is a sudoku board.
type Board struct {
	// Rows and cols are separate, but currently assumed to be the same (i.e. a square board).
	rows   int
	cols   int
	values map[Coordinate]int
	rules  Rules
}

// ParseBoard parses the given string representation of a board.
// Representation is expected to be a 9x9 newline-separated string, with each character either a space or a 1-9 digit.
// Representation may have leading/trailing new lines, as well as leading/trailing tabs within lines.
func ParseBoard(s string) Board {
	values := make(map[Coordinate]int)

	s = strings.Trim(s, "\n")
	for row, rowString := range strings.Split(s, "\n") {
		for col, digit := range strings.Trim(rowString, "\t") {
			if digit == ' ' {
				continue
			}
			// Coordinates are 1-indexed.
			values[NewCoordinate(row+1, col+1)] = int(digit - '0')
		}
	}
	return NewStandardBoard(values)
}

// NewEmptyBoard returns a standard, empty sudoku board.
func NewEmptyBoard() Board {
	initialValues := make(map[Coordinate]int)
	return NewStandardBoard(initialValues)
}

// NewStandardBoard returns a standard sudoku board, with the specified initial cells populated.
func NewStandardBoard(initialValues map[Coordinate]int) Board {
	values := make(map[Coordinate]int)
	for k, v := range initialValues {
		values[k] = v
	}

	return Board{rows: 9, cols: 9, values: values}
}

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

// GetAllCoordinates returns all coordinates in the board.
func (b Board) GetAllCoordinates() []Coordinate {
	coordinates := make([]Coordinate, 0)
	for row := 1; row <= b.rows; row++ {
		for col := 1; col <= b.cols; col++ {
			coordinates = append(coordinates, NewCoordinate(row, col))
		}
	}
	return coordinates
}

// Row returns the coordinates in the specified row.
func (b Board) Row(row int) []Coordinate {
	coordinates := make([]Coordinate, 0)
	for col := 1; col <= b.cols; col++ {
		coordinates = append(coordinates, NewCoordinate(row, col))
	}
	return coordinates
}

// Col returns the coordinates in the specified col.
func (b Board) Col(col int) []Coordinate {
	coordinates := make([]Coordinate, 0)
	for row := 1; row <= b.rows; row++ {
		coordinates = append(coordinates, NewCoordinate(row, col))
	}
	return coordinates
}

// Region returns the cells in the specified region.
func (b Board) Region(row, col int) []Coordinate {
	coordinates := make([]Coordinate, 0)
	for i := 1; i <= 3; i++ {
		r := 3*(row-1) + i
		for j := 1; j <= 3; j++ {
			c := 3*(col-1) + j
			coordinates = append(coordinates, NewCoordinate(r, c))
		}
	}
	return coordinates
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

func (b Board) String() string {
	rowStrings := make([]string, 0)
	for row := 1; row <= 9; row++ {
		rowChars := make([]string, 0)
		for col := 1; col <= 9; col++ {
			value, ok := b.values[NewCoordinate(row, col)]
			if ok {
				rowChars = append(rowChars, strconv.Itoa(value))
			} else {
				rowChars = append(rowChars, " ")
			}

			if col%3 == 0 {
				rowChars = append(rowChars, "|")
			}
		}
		rowString := strings.Join(rowChars[:len(rowChars)-1], " ")
		rowStrings = append(rowStrings, rowString)

		if row%3 == 0 {
			rowStrings = append(rowStrings, strings.Repeat("-", len(rowString)))
		}
	}
	return strings.Join(rowStrings[:len(rowStrings)-1], "\n")
}
