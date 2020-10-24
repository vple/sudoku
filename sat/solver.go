package sat

import (
	"fmt"
	"time"
)

const showIterationTimes = false

// Solve attempts to solve the given formula, given the initial state.
func Solve(formula ConjunctiveFormula, state map[string]bool) (map[string]bool, bool) {
	var start time.Time
	if showIterationTimes {
		start = time.Now()
	}

	// Evaluate and unit propagate as much as possible.
	ok := true
	for ok {
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
	if showIterationTimes {
		fmt.Println(time.Since(start))
	}
	testState[litName] = true
	if solutionState, ok := Solve(formula, testState); ok {
		return solutionState, true
	}

	testState[litName] = false
	if solutionState, ok := Solve(formula, testState); ok {
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
			changed = true
			switch literal := clause.literals[0].(type) {
			case PositiveLiteral:
				newState[literal.GetName()] = true
			case NegativeLiteral:
				newState[literal.GetName()] = false
			default:
				panic("Unexpected type!")
			}
		}
	}

	return newState, changed
}

func selectLiteral(formula ConjunctiveFormula) string {
	// TODO: Pick literal better.
	// return formula.clauses[0].literals[0].GetName()
	// return fromShortestClause(formula)
	return mostFrequentLiteral(formula)
}

func fromShortestClause(formula ConjunctiveFormula) string {
	best := formula.clauses[0]
	minLength := len(best.literals)

	for _, clause := range formula.clauses[1:] {
		length := len(clause.literals)
		if length < minLength {
			best = clause
			minLength = length
		}
	}

	return best.literals[0].GetName()
}

func mostFrequentLiteral(formula ConjunctiveFormula) string {
	frequency := make(map[string]int)
	for _, clause := range formula.clauses {
		for _, literal := range clause.literals {
			frequency[literal.GetName()] = frequency[literal.GetName()] + 1
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
