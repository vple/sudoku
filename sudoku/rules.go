package sudoku

// Rules are rules to use in addition to standard sudoku rules.
type Rules struct {
	// If true, main diagonals must contain unique values.
	diagonalsUnique bool
}

// RulesVanilla are the default sudoku rules.
var RulesVanilla = Rules{}
