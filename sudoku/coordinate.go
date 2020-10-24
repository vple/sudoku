package sudoku

import (
	"fmt"
	"regexp"
	"strconv"

	"../sat"
)

var digitRegex = regexp.MustCompile("[0-9]")

// Coordinate specifies a cell position on a board.
type Coordinate struct {
	row int
	col int
}

// Coordinates is a slice of coordinates.
type Coordinates []Coordinate

// NewCoordinate creates a new coordinate.
func NewCoordinate(row, col int) Coordinate {
	return Coordinate{row, col}
}

// ParseName parses a variable name back to a coordinate and its value.
func ParseName(name string) (Coordinate, int) {
	digits := digitRegex.FindAllString(name, -1)
	row, _ := strconv.Atoi(digits[0])
	col, _ := strconv.Atoi(digits[1])
	value, _ := strconv.Atoi(digits[2])
	return NewCoordinate(row, col), value
}

// varName is a unique variable name that represents the cell at this coordinate containing a specified value.
func (c Coordinate) varName(value int) string {
	if c.row < 1 || c.col < 1 {
		panic("wtf")
	}
	return fmt.Sprintf("(%d,%d): %d", c.row, c.col, value)
}

// Literal is the literal representing this cell containing a specified value.
func (c Coordinate) Literal(value int) sat.Literal {
	return sat.NewLiteral(c.varName(value))
}

// Literals returns literals that represent all possible states for this cell.
func (c Coordinate) Literals() []sat.Literal {
	literals := make([]sat.Literal, 0)
	for i := 1; i <= 9; i++ {
		literals = append(literals, c.Literal(i))
	}
	return literals
}

// Clauses returns the clauses specifying the constraints for this coordinate.
func (c Coordinate) Clauses() []sat.DisjunctiveClause {
	return sat.ExactlyOneTrue(c.Literals())
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.row, c.col)
}
