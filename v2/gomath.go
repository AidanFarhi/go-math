// Package gomath includes various mathimatical functions
package gomath

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two numbers together
func Add[T Number](x, y T) T {
	return x + y
}
