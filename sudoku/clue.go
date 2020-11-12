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

// Sum specifies cells that sum to a given value.
type Sum struct {
	coordinates []Coordinate
	sum         int
}

// NewSum creates a new Sum.
func NewSum(coordinates []Coordinate, sum int) Sum {
	return Sum{coordinates, sum}
}

// Sum5 creates a sum summing to 5.
func Sum5(coordinates ...Coordinate) Sum {
	return NewSum(coordinates, 5)
}

// Sum10 creates a sum summing to 10.
func Sum10(coordinates ...Coordinate) Sum {
	return NewSum(coordinates, 10)
}

// Apply applies this clue to a board, returning the corresponding constraints.
func (s Sum) Apply(board Board) []Constraint {
	return []Constraint{
		NewConstantSumConstraint(s.coordinates, s.sum),
	}
}

// LittleKiller specifies a diagonal that sums to a given value.
type LittleKiller struct {
	coordinates []Coordinate
	sum         int
}

// NewLittleKiller creates a new LittleKiller.
func NewLittleKiller(diagonal []Coordinate, sum int) LittleKiller {
	return LittleKiller{diagonal, sum}
}

// Apply applies this clue to a board, returning the corresponding constraints.
func (lk LittleKiller) Apply(board Board) []Constraint {
	return []Constraint{
		NewConstantSumConstraint(lk.coordinates, lk.sum),
	}
}

// KillerCage specifies cells that have unique values, summing to a given value.
type KillerCage struct {
	coordinates []Coordinate
	sum         int
}

// NewKillerCage creates a new KillerCage.
func NewKillerCage(coordinates []Coordinate, sum int) KillerCage {
	return KillerCage{coordinates, sum}
}

// Apply applies this clue to a board, returning the corresponding constraints.
func (kc KillerCage) Apply(board Board) []Constraint {
	return []Constraint{
		NewConstantSumConstraint(kc.coordinates, kc.sum),
		NewUniqueValueConstraint(kc.coordinates...),
	}
}
