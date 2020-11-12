package conversion

import (
	"fmt"
	"sort"

	sudoku ".."
	"../../sat"
)

// ParseState parses boolean state back into a board.
func ParseState(state map[string]bool) sudoku.Board {
	initialValues := make(map[sudoku.Coordinate]int)
	for name, v := range state {
		if v {
			coordinate, value := fromName(name)
			initialValues[coordinate] = value
		}
	}
	return sudoku.NewStandardBoard(initialValues)
}

// ToFormula converts a board to CNF form.
func ToFormula(board sudoku.Board) (formula sat.ConjunctiveFormula) {
	constraints := board.AllConstraints()

	for _, constraint := range constraints {
		formula = formula.And(convert(constraint, board))
	}

	return formula
}

func convert(c sudoku.Constraint, board sudoku.Board) sat.ConjunctiveFormula {
	switch constraint := c.(type) {
	case sudoku.CellValueConstraint:
		return constrainToValues(constraint.Coordinate(), constraint.Values())
	case sudoku.UniqueValueConstraint:
		return uniqueValues(constraint.Coordinates(), board.AllValues())
	case sudoku.ContainsValuesConstraint:
		return Appears(constraint.Coordinates(), constraint.Values()...)
	case sudoku.IncreasingValueConstraint:
		formula := sat.EmptyConjunctiveFormula()

		clauses := make([]sat.DisjunctiveClause, 0)
		coordinates := constraint.Coordinates()
		size := board.Size()
		dof := size - len(coordinates) // Degrees of freedom
		for i, a := range coordinates {
			possibleAValues := make([]int, 0)
			for x := 0; x <= dof; x++ {
				possibleAValues = append(possibleAValues, i+x+1)
			}
			formula = formula.And(constrainToValues(a, possibleAValues))

			for _, b := range coordinates[i+1:] {
				// Coordinate A < Coordinate B
				for _, aValue := range board.AllValues() {
					for bValue := 1; bValue < aValue; bValue++ {
						// For aValue > bValue, we want:
						// !aValue || !bValue
						notA := toLiteral(a, aValue).Negate()
						notB := toLiteral(b, bValue).Negate()
						clauses = append(clauses, sat.NewDisjunctiveClause(notA, notB))
					}
				}
			}
		}

		return formula.And(sat.NewConjunctiveFormula(clauses))
	case sudoku.ConstantSumConstraint:
		allValues := board.AllValues()
		literals := make([]sat.Literal, 0)
		for _, coordinate := range constraint.Coordinates() {
			literals = append(literals, toLiterals(coordinate, allValues)...)
		}

		sumLiteral := sumLiterals(literals, constraint.Sum(), len(constraint.Coordinates()), board.Size())
		return sat.NewDisjunctiveClause(sumLiteral).ToFormula()
	default:
		panic(fmt.Sprintf("Unknown constraint type: %T", c))
	}
}

func constrainToValues(coordinate sudoku.Coordinate, values []int) sat.ConjunctiveFormula {
	literals := make([]sat.Literal, 0)
	for _, value := range values {
		literals = append(literals, toLiteral(coordinate, value))
	}
	return sat.ExactlyOneTrue(literals)
}

// sumLiterals creates a composite literal that is true if there are numSummands cells with values adding to sum.
// It only considers the given literals, and assumes each summand has a value from [1, maxValue].
// Repeated summands are allowed.
func sumLiterals(l []sat.Literal, sum int, numSummands int, maxValue int) sat.CompositeLiteral {
	return sat.NewCompositeLiteral(l, func(literals []sat.Literal, state map[string]bool) interface{} {
		currentSum := 0
		currentSummands := 0
		undeterminedLiterals := make([]sat.Literal, 0)

		for _, literal := range literals {
			switch result := literal.Evaluate(state).(type) {
			case bool:
				_, isPositive := literal.(sat.PositiveLiteral)
				_, isNegative := literal.(sat.NegativeLiteral)
				hasValue := (result && isPositive) || (!result && isNegative)

				if hasValue {
					_, value := fromName(literal.Name())
					currentSum += value
					currentSummands++
				}
			case sat.Literal:
				undeterminedLiterals = append(undeterminedLiterals, result)
			default:
				panic("Unexpected result type!")
			}
		}

		// We have all the info needed to fully evaluate the sum.
		if len(undeterminedLiterals) == 0 {
			return currentSum == sum
		}

		// We've found the desired number of summands.
		if currentSummands == numSummands {
			return currentSum == sum
		}

		remainingSum := sum - currentSum
		remainingSummands := numSummands - currentSummands

		// We will always exceed the target sum with the number of summands we have left.
		// This also works for remainingSummands == 0.
		if remainingSummands*1 > remainingSum {
			return false
		}

		// It's impossible to reach the target sum with the number of summands we have left.
		if remainingSummands*maxValue < remainingSum {
			return false
		}

		// // Eliminate literals with values that are too high.
		// newUndeterminedLiterals := make([]sat.Literal, 0)
		// for _, literal := range undeterminedLiterals {
		// 	_, value := fromName(literal.Name())
		// 	if value <= remainingSum {
		// 		newUndeterminedLiterals = append(newUndeterminedLiterals, literal)
		// 	}
		// }
		// undeterminedLiterals = newUndeterminedLiterals

		// There aren't enough possible values to reach the total.
		if len(undeterminedLiterals) < remainingSummands {
			return false
		}

		// More comprehensive checks to see if we can reach the target sum.
		minValues := minLiteralValues(undeterminedLiterals, remainingSummands)
		minSum := 0
		for _, value := range minValues {
			minSum += value
		}
		if minSum > remainingSum {
			return false
		}

		maxValues := maxLiteralValues(undeterminedLiterals, remainingSummands)
		maxSum := 0
		for _, value := range maxValues {
			maxSum += value
		}
		if maxSum < remainingSum {
			return false
		}

		// If there's only one summand left, look for matching literals.
		if remainingSummands == 1 {
			newUndeterminedLiterals := make([]sat.Literal, 0)
			for _, literal := range undeterminedLiterals {
				_, value := fromName(literal.Name())
				if value == remainingSum {
					newUndeterminedLiterals = append(newUndeterminedLiterals, literal)
				}
			}

			undeterminedLiterals = newUndeterminedLiterals
			switch len(undeterminedLiterals) {
			case 0:
				// The last cell can't have the desired value.
				return false
			case 1:
				// There's exactly one possibility left, so we can reduce to a non-composite literal.
				return undeterminedLiterals[0]
			}
		}

		// if remainingSummands == 2 {
		// 	fmt.Println(remainingSum, remainingSummands, undeterminedLiterals)
		// }

		// if currentSum != 0 {
		// 	fmt.Println(len(undeterminedLiterals), sum, currentSum, remainingSummands)
		// 	fmt.Println(undeterminedLiterals)
		// }

		return sumLiterals(undeterminedLiterals, remainingSum, remainingSummands, maxValue)
	})
}

func minLiteralValues(literals []sat.Literal, n int) []int {
	values := make([]int, 0)
	for _, literal := range literals {
		_, value := fromName(literal.Name())
		values = append(values, value)
	}
	sort.Ints(values)
	return values[:n]
}

func maxLiteralValues(literals []sat.Literal, n int) []int {
	values := make([]int, 0)
	for _, literal := range literals {
		_, value := fromName(literal.Name())
		values = append(values, value)
	}
	sort.Ints(values)
	return values[len(values)-n : len(values)]
}
