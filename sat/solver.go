package sat

import (
	"fmt"
	"time"
)

const showIterationTimes = false

// Solve attempts to solve the given formula, given the initial state.
func Solve(formula ConjunctiveFormula, state map[string]bool, display func(map[string]bool) string) (map[string]bool, bool) {
	var start time.Time
	if showIterationTimes {
		start = time.Now()
	}

	fmt.Println(display(state))

	// Evaluate and unit propagate as much as possible.
	ok := true
	for ok {
		// fmt.Println(len(state))
		switch expr := formula.Evaluate(state).(type) {
		case bool:
			if expr {
				return state, true
			}
			return nil, false
		case ConjunctiveFormula:
			state, ok = propagate(expr, state)
			formula = expr
		default:
			panic("Unexpected type!")
		}
	}

	testState := make(map[string]bool)
	for k, v := range state {
		testState[k] = v
	}

	litName := selectLiteral(formula)
	fmt.Println(litName)
	if showIterationTimes {
		fmt.Println(time.Since(start))
	}

	// False first is better for visualizing. This causes earlier assumptions to stay on the board longer.
	// True first might be better for speed, since setting a value to true has a lot of downstream propagation.

	testState[litName] = true
	// fmt.Println(true)
	if solutionState, ok := Solve(formula, testState, display); ok {
		return solutionState, true
	}

	testState[litName] = false
	// fmt.Println(false)
	if solutionState, ok := Solve(formula, testState, display); ok {
		return solutionState, true
	}

	return state, false
}

// propagate uses unit propagation to determine additional variable values
func propagate(formula ConjunctiveFormula, state map[string]bool) (newState map[string]bool, changed bool) {
	newState = make(map[string]bool)
	for k, v := range state {
		newState[k] = v
	}

	for _, clause := range formula.clauses {
		if len(clause.literals) == 1 {
			switch literal := clause.literals[0].(type) {
			case PositiveLiteral:
				changed = true
				newState[literal.Name()] = true
			case NegativeLiteral:
				changed = true
				newState[literal.Name()] = false
			case CompositeLiteral:
				continue
			default:
				panic("Unexpected type!")
			}
		}
	}

	return newState, changed
}

func selectLiteral(formula ConjunctiveFormula) string {
	// TODO: Pick literal better.
	// return formula.clauses[0].literals[0].Name()
	return fromShortestClause(formula)
	// return mostFrequentLiteral(formula)
	// return mostFrequentPositiveLiteral(formula)
}

func fromShortestClause(formula ConjunctiveFormula) string {
	best := formula.clauses[0]
	// minLength := len(best.literals)
	minLength := 0
	for _, literal := range best.literals {
		minLength += len(literal.Names())
	}

	for _, clause := range formula.clauses[1:] {
		// length := len(clause.literals)
		length := 0
		for _, literal := range clause.literals {
			length += len(literal.Names())
		}
		if length < minLength {
			best = clause
			minLength = length
		}
	}

	return best.literals[0].Names()[0]
}

func mostFrequentLiteral(formula ConjunctiveFormula) string {
	frequency := make(map[string]int)
	for _, clause := range formula.clauses {
		for _, literal := range clause.literals {
			for _, name := range literal.Names() {
				frequency[name] = frequency[name] + 10
				// Bias towards positive literals.
				if _, ok := literal.(NegativeLiteral); !ok {
					frequency[name] = frequency[name] + 4
				}
			}
		}
	}

	var name string
	maxFrequency := -1
	for k, v := range frequency {
		if v > maxFrequency {
			name = k
			maxFrequency = v
		}
	}

	return name
}

func mostFrequentPositiveLiteral(formula ConjunctiveFormula) string {
	frequency := make(map[string]int)
	for _, clause := range formula.clauses {
		for _, literal := range clause.literals {
			if _, ok := literal.(NegativeLiteral); ok {
				continue
			}
			for _, name := range literal.Names() {
				frequency[name] = frequency[name] + 1
			}
		}
	}

	var name string
	maxFrequency := -1
	for k, v := range frequency {
		if v > maxFrequency {
			name = k
			maxFrequency = v
		}
	}

	return name
}
