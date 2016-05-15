// RPN Calculator
package main

import (
	"errors"
	"strconv"
)

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

// PerformOperation takes two float64 operands and one of the
// valid operators (+, -, *, /) and returns the float64 value
// of the operation. operand2 cannot be zero when performing
// division. Invalid operations return a value of 0 and raise
// an error.
func PerformOperation(operand1, operand2 float64, operator string) (result float64, err error) {
	// get result
	switch operator {
	case "+":
		result = operand1 + operand2
		break
	case "-":
		result = operand1 - operand2
		break
	case "*":
		result = operand1 * operand2
		break
	case "/":
		if operand2 == 0 {
			return 0, errors.New("Cannot divide by zero.")
		}
		result = operand1 / operand2
		break
	default:
		return 0, errors.New("Unsupported operator.")
	}

	return result, nil
}
