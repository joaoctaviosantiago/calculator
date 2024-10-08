// Package calculator does simple calculations.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding
// them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting a from b.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers a and b, and
// returns the result of multiplying them together
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers a and b, and returns
// either the result of the division of a by b or
// it reutrns an error in case the division is invalid
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}

	return a / b, nil
}

// SquareRoot takes a number "a" and returns either
// its squareroot or an error if "a" is negative
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("taking the square root of a negative number is not a valid operation")
	}

	return math.Sqrt(a), nil
}
