// Package gomath includes various mathimatical functions
package gomath

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two numbers together and returns the sum
// Uses another package.
func Add[T Number](x, y T) T {
	return x + y
}
