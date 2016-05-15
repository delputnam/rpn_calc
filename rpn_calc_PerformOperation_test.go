package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestPerformOperationSuccess(t *testing.T) {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)

	// add 1000 random float64s and check
	for tc := 0; tc < 1000; tc++ {
		operand1 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
		operand2 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
		expectedResult := operand1 + operand2
		gotResult, gotErr := PerformOperation(operand1, operand2, "+")
		if expectedResult != gotResult {
			t.Errorf("Expected %v + %v = %v, Got: %v", operand1, operand2, expectedResult, gotResult)
		}
		if gotErr != nil {
			t.Errorf("Expected no errors, Got: %v", gotErr)
		}
	}

	// subtract 1000 random float64s and check
	for tc := 0; tc < 1000; tc++ {
		operand1 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
		operand2 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
		expectedResult := operand1 - operand2
		gotResult, gotErr := PerformOperation(operand1, operand2, "-")
		if expectedResult != gotResult {
			t.Errorf("Expected %v - %v = %v, Got: %v", operand1, operand2, expectedResult, gotResult)
		}
		if gotErr != nil {
			t.Errorf("Expected no errors, Got: %v", gotErr)
		}
	}

	// multiply 1000 random float64s and check
	for tc := 0; tc < 1000; tc++ {
		operand1 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
		operand2 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
		expectedResult := operand1 * operand2
		gotResult, gotErr := PerformOperation(operand1, operand2, "*")
		if expectedResult != gotResult {
			t.Errorf("Expected %v * %v = %v, Got: %v", operand1, operand2, expectedResult, gotResult)
		}
		if gotErr != nil {
			t.Errorf("Expected no errors, Got: %v", gotErr)
		}
	}

	// divide 1000 random float64s and check
	for tc := 0; tc < 1000; tc++ {
		operand1 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
		operand2 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
		// ensure operand 2 is non-zero
		for operand2 == 0 {
			operand2 = (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
		}
		expectedResult := operand1 / operand2
		gotResult, gotErr := PerformOperation(operand1, operand2, "/")
		if expectedResult != gotResult {
			t.Errorf("Expected %v / %v = %v, Got: %v", operand1, operand2, expectedResult, gotResult)
		}
		if gotErr != nil {
			t.Errorf("Expected no errors, Got: %v", gotErr)
		}
	}
}

func TestPerformOperationFail(t *testing.T) {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)

	// divide by zero should fail
	operand1 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
	expectedResult := float64(0)
	gotResult, gotErr := PerformOperation(operand1, 0, "/")

	if expectedResult != gotResult {
		t.Errorf("Expected result of 0 when performing division by zero, Got: %v", gotResult)
	}

	if gotErr == nil {
		t.Errorf("Expected error when dividing by zero, Got nil.")
	}

	// Test all invalid operators
	for i := 0; i < len(AllInvalids); i++ {
		tc := string(AllInvalids[i])
		operand1 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
		operand2 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))

		expectedResult := float64(0)
		gotResult, gotErr := PerformOperation(operand1, operand2, tc)

		if expectedResult != gotResult {
			t.Errorf("Expected result of 0 when using invalid operator, Got: %v", gotResult)
		}

		if gotErr == nil {
			t.Errorf("Expected error when using invalid operator, Got nil.")
		}
	}

	// Test some random strings of random length greater than 2
	// since all multi-character strings are invalid.
	aLen := len(ASCII)
	for tn := 0; tn < 1000; tn++ {
		len := rand.Intn(100) + 2 // min length of 2 chars
		var tc string
		for i := 0; i < len; i++ {
			aRand := rand.Intn(aLen)
			tc += string(ASCII[aRand])
		}
		operand1 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))
		operand2 := (rnd.Float64() * float64(rnd.Intn(10000))) - float64(rnd.Intn(10000))

		expectedResult := float64(0)
		gotResult, gotErr := PerformOperation(operand1, operand2, tc)

		if expectedResult != gotResult {
			t.Errorf("Expected result of 0 when using invalid operator, Got: %v", gotResult)
		}

		if gotErr == nil {
			t.Errorf("Expected error when using invalid operator, Got nil.")
		}
	}
}
