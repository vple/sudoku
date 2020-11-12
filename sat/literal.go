package sat

import (
	"fmt"
	"strings"
)

// Expression is a boolean expression.
type Expression interface {
	// Evaluate evaluates this expression to a bool if possible, otherwise a simplified Expression of the same type.
	Evaluate(map[string]bool) interface{}
}

// Literal is a boolean value.
type Literal interface {
	Name() string
	Names() []string
	Negate() Literal
	Evaluate(map[string]bool) interface{}
}

// Literals is a slice of literals.
type Literals []Literal

// PositiveLiteral is a variable.
type PositiveLiteral struct {
	name string
}

// NewLiteral creates a new literal.
func NewLiteral(name string) PositiveLiteral {
	return PositiveLiteral{name: name}
}

// Negate negates this literal.
func (pl PositiveLiteral) Negate() Literal {
	return NegativeLiteral{pl}
}

// Evaluate returns this literal's value if known.
func (pl PositiveLiteral) Evaluate(state map[string]bool) interface{} {
	value, ok := state[pl.Name()]
	if !ok {
		return pl
	}
	return value
}

// Name returns the name of the variable in this literal.
func (pl PositiveLiteral) Name() string {
	return pl.name
}

// Names returns the names of the variables in this literal.
func (pl PositiveLiteral) Names() []string {
	return []string{pl.Name()}
}

func (pl PositiveLiteral) String() string {
	return pl.Name()
}

// NegativeLiteral is the negation of a variable.
type NegativeLiteral struct {
	literal Literal // The literal being negated.
}

// Negate negates this literal.
func (nl NegativeLiteral) Negate() Literal {
	return nl.literal
}

// Evaluate returns this literal's value if known.
func (nl NegativeLiteral) Evaluate(state map[string]bool) interface{} {
	switch result := nl.literal.Evaluate(state).(type) {
	case bool:
		return !result
	case Literal:
		return result.Negate()
	default:
		panic("Unexpected evaluation result!")
	}
}

// Name returns the name of the variable in this literal.
func (nl NegativeLiteral) Name() string {
	return nl.literal.Name()
}

// Names returns the names of the variables in this literal.
func (nl NegativeLiteral) Names() []string {
	return []string{nl.Name()}
}

func (nl NegativeLiteral) String() string {
	return fmt.Sprintf("~%s", nl.literal)
}

// A CompositeLiteral is a lazy-evaluated "literal" that asserts a condition containing multiple literals.
type CompositeLiteral struct {
	literals []Literal
	reduce   func([]Literal, map[string]bool) interface{}
}

// NewCompositeLiteral creates a new CompositeLiteral.
func NewCompositeLiteral(literals []Literal, reduce func([]Literal, map[string]bool) interface{}) CompositeLiteral {
	return CompositeLiteral{literals, reduce}
}

// Negate negates this literal.
func (cl CompositeLiteral) Negate() Literal {
	return NegativeLiteral{cl}
}

// Evaluate returns this literal's value if known.
func (cl CompositeLiteral) Evaluate(state map[string]bool) interface{} {
	return cl.reduce(cl.literals, state)
}

// Name returns the name of the variable in this literal.
func (cl CompositeLiteral) Name() string {
	return fmt.Sprintf("Composite[%s]", strings.Join(cl.Names(), ","))
}

// Names returns the names of the variables in this literal.
func (cl CompositeLiteral) Names() []string {
	names := make([]string, 0)
	for _, literal := range cl.literals {
		names = append(names, literal.Name())
	}
	return names
}
