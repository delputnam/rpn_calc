// RPN Calculator
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/delputnam/rpn_calc/stack"
)

func main() {

	var s stack.Stack
	var input string
	var output float64

	// open Stdin reader
	reader := bufio.NewReader(os.Stdin)

	// begin loop
	for {

		// prompt for input
		fmt.Print("> ")

		// read until new line
		// Note: I'm ignoring the isPrefix return value here since legal operands
		// are all shorter than MaxScanTokenSize of 65,536
		bytes, _, err := reader.ReadLine()

		// exit if read an EOF
		if err == io.EOF {
			os.Exit(0)
		}

		// convert read bytes to string
		input = string(bytes)

		// exit if input is "q"
		if input == "q" {
			os.Exit(0)
		}

		// if valid operand, push onto stack and echo to console
		num, err := IsOperand(input)
		if err == nil {

			// push num onto stack
			s.Push(num)

			// display output and continue with loop
			fmt.Println(num)
			continue
		}

		// if the input is aa valid operator, perform operation, push result onto
		// stack and echo to console
		operator, err := IsOperator(input)
		if err == nil {

			// all of the supported operators require exactly 2 operands, so
			// make sure we have at least 2 operands in stack to continue
			if s.Size() < 2 {
				fmt.Println("Not enough operands. Ignoring operator.")
				continue
			}

			// get operands. note that the order of the operands is important
			// the first one off the stack is the second operand and the second
			// one off the stack is the first operand.
			operand2, err := s.Pop()
			if err != nil {
				log.Fatalf("Error getting operand from stack: %v", err)
			}
			operand1, err := s.Pop()
			if err != nil {
				log.Fatalf("Error getting operand from stack: %v", err)
			}

			// get result
			output, err = PerformOperation(operand1, operand2, operator)
			if err != nil {
				fmt.Printf("Error performing operation '%v': %v Ignoring.\n", operator, err)
				// operation failed, so push operators back onto stack and continue
				s.Push(operand1)
				s.Push(operand2)
				continue
			}

			// push result onto stack
			s.Push(output)

			// display result and continue with loop
			fmt.Println(output)
			continue
		}

		// if we're here, the input was not a valid operator or operand
		fmt.Println("Invalid input. Ignoring.")
	}
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
