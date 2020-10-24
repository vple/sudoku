package sudoku

import (
	"strconv"
	"strings"
)

// Board is a sudoku board.
type Board struct {
	// Board is assumed to be square.
	size   int
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

	return Board{size: 9, values: values}
}

// AllCoordinates returns all coordinates in the board.
func (b Board) AllCoordinates() []Coordinate {
	coordinates := make([]Coordinate, 0)
	for row := 1; row <= b.size; row++ {
		for col := 1; col <= b.size; col++ {
			coordinates = append(coordinates, NewCoordinate(row, col))
		}
	}
	return coordinates
}

// Row returns the coordinates in the specified row.
func (b Board) Row(row int) []Coordinate {
	coordinates := make([]Coordinate, 0)
	for col := 1; col <= b.size; col++ {
		coordinates = append(coordinates, NewCoordinate(row, col))
	}
	return coordinates
}

// Col returns the coordinates in the specified col.
func (b Board) Col(col int) []Coordinate {
	coordinates := make([]Coordinate, 0)
	for row := 1; row <= b.size; row++ {
		coordinates = append(coordinates, NewCoordinate(row, col))
	}
	return coordinates
}

// Region returns the cells in the specified region.
func (b Board) Region(row, col int) []Coordinate {
	// Currently assumes size = 9
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
