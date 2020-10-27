package sat

import (
	"fmt"
	"strings"
)

// Clause is a conjunctive or disjunctive clause.
type Clause interface {
}

// DisjunctiveClause is a clause whose literals are ORed together.
type DisjunctiveClause struct {
	literals []Literal
}

// NewDisjunctiveClause creates a new disjunctive clause.
func NewDisjunctiveClause(literals ...Literal) DisjunctiveClause {
	return DisjunctiveClause{literals}
}

// Or returns the disjunctive clause that results from ORing these clauses together.
func (c DisjunctiveClause) Or(others ...DisjunctiveClause) DisjunctiveClause {
	literals := c.literals
	for _, other := range others {
		literals = append(literals, other.literals...)
	}
	return DisjunctiveClause{literals}
}

// Evaluate evaluates this clause, returning a simplified clause or a bool.
func (c DisjunctiveClause) Evaluate(state map[string]bool) interface{} {
	if len(c.literals) == 0 {
		// OR() = false
		return false
	}

	remainingLiterals := make(map[Literal]bool, 0)

	for _, literal := range c.literals {
		// value := literal.Evaluate(state)
		switch value := literal.Evaluate(state).(type) {
		case Literal:
			if _, ok := remainingLiterals[value]; ok {
				// OR(x, x, ...) = OR(x, ...)
				continue
			}
			if _, ok := remainingLiterals[value.Negate()]; ok {
				// OR(x, ~x, ...) = true
				return true
			}
			remainingLiterals[value] = true
		case bool:
			if value {
				return true
			}
			// We can drop the literal here.
		default:
			panic("Unexpected type!")
		}
	}

	if len(remainingLiterals) == 0 {
		return false
	}

	literals := make([]Literal, 0)
	for k := range remainingLiterals {
		literals = append(literals, k)
	}
	return DisjunctiveClause{literals}
}

// ToFormula returns a formula containing this clause.
func (c DisjunctiveClause) ToFormula() ConjunctiveFormula {
	return NewConjunctiveFormula([]DisjunctiveClause{c})
}

func (c DisjunctiveClause) String() string {
	strs := make([]string, 0)
	for _, literal := range c.literals {
		strs = append(strs, fmt.Sprintf("%s", literal))
	}
	return strings.Join(strs, " v ")
}

// ConjunctiveClause is a clause whose literals are ANDed together.
type ConjunctiveClause struct {
	literals []Literal
}

// NewConjunctiveClause creates a new conjunctive clause.
func NewConjunctiveClause(literals ...Literal) ConjunctiveClause {
	return ConjunctiveClause{literals}
}

// Evaluate evaluates this clause, returning a simplified clause or a bool.
func (c ConjunctiveClause) Evaluate(state map[string]bool) interface{} {
	if len(c.literals) == 0 {
		// AND() = true
		return true
	}

	remainingLiterals := make(map[Literal]bool, 0)

	for _, literal := range c.literals {
		// value := literal.Evaluate(state)
		switch value := literal.Evaluate(state).(type) {
		case Literal:
			if _, ok := remainingLiterals[value]; ok {
				// AND(x, x, ...) = AND(x, ...)
				continue
			}
			if _, ok := remainingLiterals[value.Negate()]; ok {
				// AND(x, ~x, ...) = false
				return false
			}
			remainingLiterals[value] = true
		case bool:
			if !value {
				// AND(false, ...) = false
				return false
			}
			// AND(true, ...) = AND(...)
			// We can drop the literal here.
		default:
			panic("Unexpected type!")
		}
	}

	if len(remainingLiterals) == 0 {
		return true
	}

	literals := make([]Literal, 0)
	for k := range remainingLiterals {
		literals = append(literals, k)
	}
	return ConjunctiveClause{literals}
}

// ToCNF returns this clause in conjunctive normal form.
func (c ConjunctiveClause) ToCNF() ConjunctiveFormula {
	clauses := make([]DisjunctiveClause, 0)
	for _, literal := range c.literals {
		clauses = append(clauses, NewDisjunctiveClause(literal))
	}
	return ConjunctiveFormula{clauses}
}

func (c ConjunctiveClause) String() string {
	strs := make([]string, 0)
	for _, literal := range c.literals {
		strs = append(strs, fmt.Sprintf("%s", literal))
	}
	return strings.Join(strs, " ^ ")
}
