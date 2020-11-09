package sudoku

import (
	"fmt"
	"strconv"
	"strings"
)

// Board is a sudoku board.
type Board struct {
	// Board is assumed to be square.
	size         int
	regionHeight int // The number of rows in a region.
	regionWidth  int // The number of cols in a region.
	values       map[Coordinate]int
	constraints  []Constraint // The constraints specifying the current board state. Does not have to include constraints for cell values.
	rules        []Rule
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

	rules := []Rule{BasicSudokuRules{}}

	return Board{size: 9, regionHeight: 3, regionWidth: 3, values: values, rules: rules}
}

// AddRules adds the specified rules to this board.
func (b Board) AddRules(rules ...Rule) {
	b.rules = append(b.rules, rules...)
}

// Size is the size of this board.
func (b Board) Size() int {
	return b.size
}

// InBounds returns whether or not the given coordinate is in the bounds of this board.
func (b Board) InBounds(coordinate Coordinate) bool {
	if coordinate.Row() < 1 || coordinate.Row() > b.size {
		return false
	}
	if coordinate.Col() < 1 || coordinate.Col() > b.size {
		return false
	}
	return true
}

// AllValues returns the possible values on the board.
func (b Board) AllValues() []int {
	values := make([]int, 0)
	for i := 1; i <= b.size; i++ {
		values = append(values, i)
	}

	return values
}

// Value returns the value at this coordinate in the board.
func (b Board) Value(c Coordinate) (int, bool) {
	value, ok := b.values[c]
	return value, ok
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

// AllRows returns all rows in the board.
func (b Board) AllRows() [][]Coordinate {
	rows := make([][]Coordinate, 0)
	for i := 1; i <= b.size; i++ {
		rows = append(rows, b.Row(i))
	}
	return rows
}

// Col returns the coordinates in the specified col.
func (b Board) Col(col int) []Coordinate {
	coordinates := make([]Coordinate, 0)
	for row := 1; row <= b.size; row++ {
		coordinates = append(coordinates, NewCoordinate(row, col))
	}
	return coordinates
}

// AllCols returns all cols in the board.
func (b Board) AllCols() [][]Coordinate {
	cols := make([][]Coordinate, 0)
	for i := 1; i <= b.size; i++ {
		cols = append(cols, b.Col(i))
	}
	return cols
}

// Region returns the cells in the specified region.
func (b Board) Region(regionRow, regionCol int) []Coordinate {
	coordinates := make([]Coordinate, 0)
	for i := 1; i <= b.regionHeight; i++ {
		row := (regionRow-1)*b.regionHeight + i
		for j := 1; j <= b.regionWidth; j++ {
			col := (regionCol-1)*b.regionWidth + j
			coordinates = append(coordinates, NewCoordinate(row, col))
		}
	}

	return coordinates
}

// AllRegions returns all regions in the board.
func (b Board) AllRegions() [][]Coordinate {
	regions := make([][]Coordinate, 0)
	for row := 1; row <= b.size/b.regionHeight; row++ {
		for col := 1; col <= b.size/b.regionWidth; col++ {
			regions = append(regions, b.Region(row, col))
		}
	}

	return regions
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

// KnightMoves returns all coordinates that are a knight's move away from the given coordinate.
func (b Board) KnightMoves(coordinate Coordinate) (coordinates Coordinates) {
	long := []int{2, -2}
	short := []int{1, -1}

	row := coordinate.Row()
	col := coordinate.Col()

	for _, l := range long {
		for _, s := range short {
			c1 := NewCoordinate(row+l, col+s)
			if b.InBounds(c1) {
				coordinates = append(coordinates, c1)
			}

			c2 := NewCoordinate(row+s, col+l)
			if b.InBounds(c2) {
				coordinates = append(coordinates, c2)
			}
		}
	}

	return coordinates
}

// AllConstraints returns the constraints from the board's rules and its initial values.
func (b Board) AllConstraints() (constraints []Constraint) {
	for _, rule := range b.rules {
		constraints = append(constraints, rule.Apply(b)...)
	}

	constraints = append(constraints, b.constraints...)

	for coordinate, value := range b.values {
		constraints = append(constraints, NewCellValueConstraint(coordinate, value))
	}

	return constraints
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
