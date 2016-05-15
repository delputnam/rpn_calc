package main

import (
	"log"
	"math/rand"
	"testing"
)

func TestStack(t *testing.T) {
	var s Stack
	var n []float64

	// push some numbers on the stack and check size of stack
	// and make sure the number just pushed is the one on TestIsOperatorSuccess
	// of the stack.
	for tc := 0; tc < 1000; tc++ {

		// generate a random float64 and remember it
		expectedNum := rand.Float64() * float64(rand.Intn(10000))
		n = append(n, expectedNum)

		// push onto stack
		s.Push(n[tc])

		// inspect stack internals to see if it was pushed correctly
		gotNum := s.s[tc]

		if expectedNum != gotNum {
			log.Fatalf("Pushed %f onto stack, saw %f as last num on stack", expectedNum, gotNum)
		}

		expectedSize := tc + 1
		gotSize := s.Size()
		if expectedSize != gotSize {
			log.Fatalf("Expected stack size %v after push, Got: %v", expectedSize, gotSize)
		}
	}

	// pop everything off the stack and make sure they match
	// and make sure the size is as expected after the pop.
	for tc := 999; tc >= 0; tc-- {
		gotNum, gotErr := s.Pop()
		expectedNum := n[tc]
		if expectedNum != gotNum {
			log.Fatalf("Expected to pop %f, Got: %f at index %v", expectedNum, gotNum, tc)
		}
		if gotErr != nil {
			log.Fatalf("Expected no error after pop, Got: %v", gotErr)
		}

		expectedSize := tc
		gotSize := s.Size()
		if expectedSize != gotSize {
			log.Fatalf("Expected size %v after pop, Got: %v", expectedSize, gotSize)
		}
	}

	// Check to make sure stack is now empty
	gotSize := s.Size()
	if gotSize != 0 {
		log.Fatalf("Expected empty stack, Got size: %v", gotSize)
	}

	// Try to pop on more to test for empty stack error
	expectedNum := float64(0)
	gotNum, gotErr := s.Pop()

	if gotErr == nil {
		log.Fatalf("Expected error when popping from empty stack, Got: nil")
	}
	if expectedNum != gotNum {
		log.Fatalf("Expected 0 when popping from empty stack, Got: %v", gotNum)
	}
}
