// RPN Calculator
package main

import (
	"errors"
	"strconv"
)

// Stack is just simple slice of float64
type Stack struct {
	s []float64
}

// Push a new number onto the stack no return value
func (s *Stack) Push(num float64) {
	s.s = append(s.s, num)
}

// Pop a number off the stack if one exists
// if stack is empty return error.
func (s *Stack) Pop() (float64, error) {
	if len(s.s) > 0 {
		num := s.s[len(s.s)-1]
		s.s = s.s[:len(s.s)-1]
		return num, nil
	}
	return 0, errors.New("Can't pop from empty stack.")
}

// Size returns the count of elements on the stack
func (s *Stack) Size() int {
	return len(s.s)
}

func main() {

	// open Stdin reader

	// begin loop

	// prompt for input

	// read until new line

	// exit on EOF

	// convert read bytes to string

	// push onto stack if valid operand

	// perform op if valid operator

	// Exit if "q"

	// check for at least 2 operands in stack to continue

	// get result

	// pop operands off stack

	// push result onto stack

	// display output
}

// IsOperand uses ParseFloat and returns a float64
// if the input is a valid representation of a number,
// otherwise it returns 0 and an error.
func IsOperand(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

// IsOperator makes sure that the input is one of the valid
// non-numeric input values. If so, it returns the input,
// otherwise it returns an empty string and an error.
func IsOperator(input string) (string, error) {
	switch input {
	case "+", "-", "*", "/":
		return input, nil
	}
	return "", errors.New("Invalid Operator")
}
