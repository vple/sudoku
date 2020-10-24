package sudoku

import (
	"fmt"
)

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

// Row is this coordinate's row.
func (c Coordinate) Row() int { return c.row }

// Col is this coordinate's column.
func (c Coordinate) Col() int { return c.col }

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.row, c.col)
}
