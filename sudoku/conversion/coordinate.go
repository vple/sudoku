package conversion

import (
	"fmt"
	"strconv"

	"regexp"

	sudoku ".."
	"../../sat"
)

var digitRegex = regexp.MustCompile("[0-9]")

func litName(coordinate sudoku.Coordinate, value int) string {
	return fmt.Sprintf("%d-%d:%d", coordinate.Row(), coordinate.Col(), value)
}

func toLiteral(coordinate sudoku.Coordinate, value int) sat.Literal {
	name := litName(coordinate, value)
	return sat.NewLiteral(name)
}

// fromName parses a variable name back to a coordinate and its value.
func fromName(name string) (sudoku.Coordinate, int) {
	digits := digitRegex.FindAllString(name, -1)
	row, _ := strconv.Atoi(digits[0])
	col, _ := strconv.Atoi(digits[1])
	value, _ := strconv.Atoi(digits[2])
	return sudoku.NewCoordinate(row, col), value
}

// toLiterals returns literals that represent all possible states for this cell.
func toLiterals(c sudoku.Coordinate) []sat.Literal {
	literals := make([]sat.Literal, 0)
	for i := 1; i <= 9; i++ {
		literals = append(literals, toLiteral(c, i))
	}
	return literals
}
