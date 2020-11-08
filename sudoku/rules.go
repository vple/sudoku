package sudoku

// Rules are rules to use in addition to standard sudoku rules.
type Rules struct {
	// If true, main diagonals must contain unique values.
	diagonalsUnique bool
}

// RulesVanilla are the default sudoku rules.
var RulesVanilla = Rules{}

// Rule specifies a rule to use when solving.
type Rule interface {
	// Apply applies this rule to a board, returning the corresponding constraints.
	Apply(Board) []Constraint
}

// BasicSudokuRules specify the default sudoku rules.
type BasicSudokuRules struct{}

// Apply applies this rule to a board, returning the corresponding constraints.
func (b BasicSudokuRules) Apply(board Board) (constraints []Constraint) {
	for _, coordinate := range board.AllCoordinates() {
		constraints = append(constraints, NewCellValueConstraint(coordinate, board.AllValues()...))
	}

	for _, row := range board.AllRows() {
		constraints = append(constraints, NewUniqueValueConstraint(row...))
	}
	for _, col := range board.AllCols() {
		constraints = append(constraints, NewUniqueValueConstraint(col...))
	}
	for _, region := range board.AllRegions() {
		constraints = append(constraints, NewUniqueValueConstraint(region...))
	}

	return constraints
}

// AntiKnightMoveRule specifies that any two cells that are a knight's move apart may not contain the same value.
type AntiKnightMoveRule struct{}

// Apply applies this rule to a board, returning the corresponding constraints.
func (a AntiKnightMoveRule) Apply(board Board) (constraints []Constraint) {
	for _, coordinate := range board.AllCoordinates() {
		for _, knight := range board.KnightMoves(coordinate) {
			constraints = append(constraints, NewUniqueValueConstraint(coordinate, knight))
		}
	}

	return constraints
}
