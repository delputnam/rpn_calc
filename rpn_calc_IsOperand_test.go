package main

import "testing"

func TestIsOperandSuccess(t *testing.T) {

	// success
	testCases := []struct {
		desc        string
		input       string
		expectedNum float64
	}{
		{
			"Convert simple integer as string to float64",
			"1",
			1,
		},
		{
			"Convert simple decimal as string to float64",
			"1.5",
			1.5,
		},
		{
			"Convert simple negative integer as string to float64",
			"-1",
			-1,
		},
		{
			"Convert simple negative decimal as string to float64",
			"-1.5",
			-1.5,
		},
		{
			"Convert MaxFloat64 as string to float64",
			"1.797693134862315708145274237317043567981e+308",
			1.797693134862315708145274237317043567981e+308,
		},
		{
			"Convert SmallestNonzeroFloat64 as string to float64",
			"4.940656458412465441765687928682213723651e-324",
			4.940656458412465441765687928682213723651e-324,
		},
	}

	for _, tc := range testCases {

		gotNum, gotErr := IsOperand(tc.input)

		if gotNum != tc.expectedNum {
			t.Errorf("While testing %s: Expected: %v, Got: %v", tc.desc, tc.expectedNum, gotNum)
		}
		if gotErr != nil {
			t.Errorf("While testing %s: %v", tc.desc, gotErr)
		}
	}

}

func TestIsOperandFail(t *testing.T) {

	// fail
	testCases := []struct {
		desc        string
		input       string
		expectedNum float64
	}{
		{
			"Attempt to convert non-numeric string to float",
			"abc",
			0,
		},
	}

	for _, tc := range testCases {

		gotNum, gotErr := IsOperand(tc.input)

		if gotNum != tc.expectedNum {
			t.Errorf("While testing %s: Expected: %v, Got: %v", tc.desc, tc.expectedNum, gotNum)
		}
		if gotErr == nil {
			t.Errorf("While testing %s, error should have been received", tc.desc)
		}
	}

}
