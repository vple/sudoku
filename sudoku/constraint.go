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

// Coordinate is the coordinate being constrained.
func (c CellValueConstraint) Coordinate() Coordinate {
	return c.coordinate
}

// Values returns the values this coordinate can contain.
func (c CellValueConstraint) Values() []int {
	return c.values
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

// Coordinates are the coordinates that must all be unique.
func (u UniqueValueConstraint) Coordinates() []Coordinate {
	return u.coordinates
}

// ContainsValuesConstraint specifies that at least one of its coordinates has each of the specified values.
type ContainsValuesConstraint struct {
	coordinates []Coordinate
	values      []int
}

// NewContainsValuesConstraint creates a new contains values constraint.
func NewContainsValuesConstraint(coordinates []Coordinate, values []int) ContainsValuesConstraint {
	return ContainsValuesConstraint{coordinates, values}
}

// Coordinates subject to the constraint.
func (c ContainsValuesConstraint) Coordinates() []Coordinate {
	return c.coordinates
}

// Values that must appear in the constrained coordinates.
func (c ContainsValuesConstraint) Values() []int {
	return c.values
}

// IncreasingValueConstraint specifies that the values in its coordinates are in strictly increasing order.
type IncreasingValueConstraint struct {
	coordinates []Coordinate
}

// NewIncreasingValueConstraint creates a new increasing value constraint.
func NewIncreasingValueConstraint(coordinates ...Coordinate) IncreasingValueConstraint {
	return IncreasingValueConstraint{coordinates}
}

// Coordinates subject to the constraint.
func (i IncreasingValueConstraint) Coordinates() []Coordinate {
	return i.coordinates
}

// ConstantSumConstraint specifies that the given cells sum to the specified constant.
type ConstantSumConstraint struct {
	coordinates []Coordinate
	sum         int
}

// NewConstantSumConstraint creates a new ConstantSumConstraint.
func NewConstantSumConstraint(coordinates []Coordinate, sum int) ConstantSumConstraint {
	return ConstantSumConstraint{coordinates, sum}
}

// Coordinates subject to the constraint.
func (c ConstantSumConstraint) Coordinates() []Coordinate {
	return c.coordinates
}

// Sum is the sum of the coordinates.
func (c ConstantSumConstraint) Sum() int {
	return c.sum
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
