package conversion

import (
	"fmt"

	sudoku ".."
)

var digitRegex = regexp.MustCompile("\d")

func litName(coordinate sudoku.Coordinate, value int) string {
	return fmt.Sprintf("%d-%d:%d", coordinate.Row(), coordinate.Col(), value)
}

func literal(coordinate sudoku.Coordinate, value int) sat.Literal {
	name := litName(coordinate, value)
	return sat.NewLiteral(name)
}

// Literals returns literals that represent all possible states for this cell.
func (c Coordinate) Literals() []sat.Literal {
	literals := make([]sat.Literal, 0)
	for i := 1; i <= 9; i++ {
		literals = append(literals, c.Literal(i))
	}
	return literals
}

// ParseName parses a variable name back to a coordinate and its value.
func ParseName(name string) (Coordinate, int) {
	digits := digitRegex.FindAllString(name, -1)
	row, _ := strconv.Atoi(digits[0])
	col, _ := strconv.Atoi(digits[1])
	value, _ := strconv.Atoi(digits[2])
	return NewCoordinate(row, col), value
}