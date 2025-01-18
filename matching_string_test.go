package main

import (
	"os"
	"testing"
)

func TestFindMatchingStrings(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedOutput string
	}{
		{
			input:          []string{"4", "abcd", "acbd", "aaab", "acbd"},
			expectedOutput: "2 4",
		},
		{
			input:          []string{"11", "Satu", "Sate", "Tujuh", "Tusuk", "Tujuh", "Sate", "Bonus", "Tiga", "Puluh", "Tujuh", "Tusuk"},
			expectedOutput: "3 5 10",
		},
		{
			input:          []string{"5", "pisang", "goreng", "enak", "sekali", "rasanya"},
			expectedOutput: "false",
		},
	}

	for idx, testCase := range testCases {
		inputReader, inputWriter, _ := os.Pipe()
		os.Stdin = inputReader

		for _, line := range testCase.input {
			inputWriter.WriteString(line + "\n")
		}
		inputWriter.Close()

		result := findMatchingStrings()
		if result != testCase.expectedOutput {
			t.Errorf("Test case %d failed: expected '%s', got '%s'", idx+1, testCase.expectedOutput, result)
		}
	}
}
