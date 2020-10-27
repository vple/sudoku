package sat

import (
	"fmt"
)

// Expression is a boolean expression.
type Expression interface {
	// Evaluate evaluates this expression to a bool if possible, otherwise a simplified Expression of the same type.
	Evaluate(map[string]bool) interface{}
}

// Literal is a boolean value.
type Literal interface {
	Name() string
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

func (pl PositiveLiteral) String() string {
	return pl.Name()
}

// NegativeLiteral is the negation of a variable.
type NegativeLiteral struct {
	variable PositiveLiteral
}

// Negate negates this literal.
func (nl NegativeLiteral) Negate() Literal {
	return nl.variable
}

// Evaluate returns this literal's value if known.
func (nl NegativeLiteral) Evaluate(state map[string]bool) interface{} {
	value, ok := state[nl.Name()]
	if !ok {
		return nl
	}
	return !value
}

// Name returns the name of the variable in this literal.
func (nl NegativeLiteral) Name() string {
	return nl.variable.Name()
}

func (nl NegativeLiteral) String() string {
	return fmt.Sprintf("~%s", nl.variable)
}
