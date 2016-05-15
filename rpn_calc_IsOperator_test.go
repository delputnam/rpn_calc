package main

import (
	"log"
	"math/rand"
	"testing"
)

const (
	Alphabet           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numerals           = "0123456789"
	InvalidNonAlphaNum = "~!@#$%^&()_={}[]\\|<,>.?\"';:`"
	AllInvalids        = Alphabet + Numerals + InvalidNonAlphaNum
	AllValids          = "+-*/"
	ASCII              = Alphabet + Numerals + InvalidNonAlphaNum + AllValids
)

func TestIsOperatorSuccess(t *testing.T) {

	for i := 0; i < len(AllValids); i++ {

		tc := string(AllValids[i])

		gotOp, gotErr := IsOperator(tc)

		if gotOp != string(AllValids[i]) {
			log.Fatalf("While testing IsOperator(\"%v\"), Expected: %v, Got: %v", tc, AllValids[i], gotOp)
		}
		if gotErr != nil {
			log.Fatalf("While testing IsOperator(\"%v\"), Expected: no error, Got %v.", tc, gotErr)
		}
	}
}

func TestIsOperatorFail(t *testing.T) {

	// Test all invalid single characters
	for i := 0; i < len(AllInvalids); i++ {
		tc := string(AllInvalids[i])

		gotOp, gotErr := IsOperator(tc)

		if gotOp != "" {
			log.Fatalf("While testing IsOperator(\"%v\"), Expected: empty string, Got: %v", tc, gotOp)
		}
		if gotErr == nil {
			log.Fatalf("While testing IsOperator(\"%v\"), error should have been received, got nil.", tc)
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
		gotOp, gotErr := IsOperator(tc)

		if gotOp != "" {
			log.Fatalf("While testing IsOperator(\"%v\"), Expected: empty string, Got: %v", tc, gotOp)
		}
		if gotErr == nil {
			log.Fatalf("While testing IsOperator(\"%v\"), error should have been received, got nil.", tc)
		}
	}
}
