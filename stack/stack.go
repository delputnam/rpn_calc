package stack

import "errors"

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
