package sudoku

// Clue is an initial constraint on the board, other than actual cell values.
type Clue interface {
	// Apply applies this clue to a board, returning the corresponding constraints.
	Apply(Board) []Constraint
}

// Thermometer indicates a consecutive sequence of cells that have strictly increasing values.
type Thermometer struct {
	coordinates []Coordinate
}

// NewThermometer creates a new thermometer.
func NewThermometer(coordinates ...Coordinate) Thermometer {
	return Thermometer{coordinates}
}

// Apply applies this clue to a board, returning the corresponding constraints.
func (t Thermometer) Apply(board Board) []Constraint {
	return []Constraint{
		NewIncreasingValueConstraint(t.coordinates...),
	}
}
