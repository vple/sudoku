package sudoku

// Constraint specifies a restriction o valid solutions.
type Constraint interface {
}

// CellValueConstraint specifies that the specified cell must contain exactly one of the given values.
type CellValueConstraint struct {
	coordinate Coordinate
	values     []int
}

// NewCellValueConstraint creates a new CellValueConstraint.
func NewCellValueConstraint(coordinate Coordinate, values ...int) CellValueConstraint {
	return CellValueConstraint{coordinate: coordinate, values: values}
}

// UniqueValueConstraint specifies that its coordinates all have unique values.
// No two coordinates have the same value.
type UniqueValueConstraint struct {
	coordinates []Coordinate
}

// NewUniqueValueConstraint creates a new unique constraint.
func NewUniqueValueConstraint(coordinates ...Coordinate) UniqueValueConstraint {
	return UniqueValueConstraint{coordinates}
}

// SumConstraint specifies that each of the specified sums are equal.
type SumConstraint struct {
	sums []summable
}

type summable interface {
}

// CellSum represents the sum of the values of the specified cells.
type CellSum struct {
	coordinates []Coordinate
}

// ConstantSum is a constant value.
type ConstantSum struct {
	value int
}
