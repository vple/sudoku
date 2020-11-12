package sat

import (
	"fmt"
	"strings"
)

// ConjunctiveFormula represents a boolean formula in conjunctive normal form.
type ConjunctiveFormula struct {
	clauses []DisjunctiveClause
}

// EmptyConjunctiveFormula returns an empty conjunctive formula.
// This evaluates to true.
func EmptyConjunctiveFormula() ConjunctiveFormula {
	clauses := make([]DisjunctiveClause, 0)
	return ConjunctiveFormula{clauses}
}

// NewConjunctiveFormula creates a formula with the given clauses.
func NewConjunctiveFormula(clauses []DisjunctiveClause) ConjunctiveFormula {
	return ConjunctiveFormula{clauses: clauses}
}

// Evaluate evaluates this formula, returning a simplified formula or a bool.
func (f ConjunctiveFormula) Evaluate(state map[string]bool) interface{} {
	remainingClauses := make([]DisjunctiveClause, 0)
	for _, clause := range f.clauses {
		switch value := clause.Evaluate(state).(type) {
		case DisjunctiveClause:
			remainingClauses = append(remainingClauses, value)
		case bool:
			if !value {
				// AND(false, ...) = false
				return false
			}
			// We can drop the clause here.
		default:
			panic("Unexpected type!")
		}
	}

	if len(remainingClauses) == 0 {
		return true
	}

	return ConjunctiveFormula{remainingClauses}
}

// And returns a formula that also contains the clauses in the other formula.
func (f ConjunctiveFormula) And(other ConjunctiveFormula) ConjunctiveFormula {
	return ConjunctiveFormula{append(f.clauses, other.clauses...)}
}

// Or returns this formula ORed with the other formulas, in CNF.
func (f ConjunctiveFormula) Or(others ...ConjunctiveFormula) ConjunctiveFormula {
	if len(others) == 0 {
		return f
	}

	other := others[0]
	clauses := make([]DisjunctiveClause, 0)
	for _, fClause := range f.clauses {
		for _, oClause := range other.clauses {
			clause := fClause.Or(oClause)
			clauses = append(clauses, clause)
		}
	}

	return ConjunctiveFormula{clauses}.Or(others[1:]...)
}

func (f ConjunctiveFormula) String() string {
	strs := make([]string, 0)
	for _, clause := range f.clauses {
		strs = append(strs, fmt.Sprintf("(%s)", clause))
	}
	return strings.Join(strs, " ^ ")
}

// DisjunctiveFormula represents a boolean formula in disjunctive normal form.
type DisjunctiveFormula struct {
	clauses []ConjunctiveClause
}

// NewDisjunctiveFormula creates a formula with the given clauses.
func NewDisjunctiveFormula(clauses []ConjunctiveClause) DisjunctiveFormula {
	return DisjunctiveFormula{clauses: clauses}
}

// Evaluate evaluates this formula, returning a simplified formula or a bool.
func (f DisjunctiveFormula) Evaluate(state map[string]bool) interface{} {
	remainingClauses := make([]ConjunctiveClause, 0)
	for _, clause := range f.clauses {
		switch value := clause.Evaluate(state).(type) {
		case ConjunctiveClause:
			remainingClauses = append(remainingClauses, value)
		case bool:
			if !value {
				// AND(false, ...) = false
				return false
			}
			// We can drop the clause here.
		default:
			panic("Unexpected type!")
		}
	}

	if len(remainingClauses) == 0 {
		return true
	}

	return DisjunctiveFormula{remainingClauses}
}

// Or returns a formula that also contains the clauses in the other formula.
func (f DisjunctiveFormula) Or(other DisjunctiveFormula) DisjunctiveFormula {
	return DisjunctiveFormula{append(f.clauses, other.clauses...)}
}

// ToCNF converts this formula to conjunctive normal form.
func (f DisjunctiveFormula) ToCNF() ConjunctiveFormula {
	if len(f.clauses) == 0 {
		return ConjunctiveFormula{[]DisjunctiveClause{NewDisjunctiveClause()}}
	}

	formulas := make([]ConjunctiveFormula, 0)
	for _, clause := range f.clauses {
		formulas = append(formulas, clause.ToCNF())
	}

	return formulas[0].Or(formulas[1:]...)
}

func (f DisjunctiveFormula) String() string {
	strs := make([]string, 0)
	for _, clause := range f.clauses {
		strs = append(strs, fmt.Sprintf("(%s)", clause))
	}
	return strings.Join(strs, " v ")
}
