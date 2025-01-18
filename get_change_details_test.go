package main

import (
	"reflect"
	"testing"
)

func TestGetChangeDetails(t *testing.T) {
	tests := []struct {
		total, paid       int
		expectedBreakdown map[int]int
		expectedChange    int
		expectedResult    bool
	}{
		{
			total:             700649,
			paid:              800000,
			expectedBreakdown: map[int]int{50000: 1, 20000: 2, 5000: 1, 2000: 2, 200: 1, 100: 1},
			expectedChange:    99300,
			expectedResult:    true,
		},
		{
			total:             575650,
			paid:              580000,
			expectedBreakdown: map[int]int{2000: 2, 200: 1, 100: 1},
			expectedChange:    4300,
			expectedResult:    true,
		},
		{
			total:             657650,
			paid:              600000,
			expectedBreakdown: nil,
			expectedChange:    0,
			expectedResult:    false,
		},
	}

	for _, test := range tests {
		breakdown, change, result := GetChangeDetails(test.total, test.paid)

		if !reflect.DeepEqual(breakdown, test.expectedBreakdown) || change != test.expectedChange || result != test.expectedResult {
			t.Errorf("For total=%d, paid=%d, expected breakdown=%v, change=%d, result=%v, got breakdown=%v, change=%d, result=%v",
				test.total, test.paid, test.expectedBreakdown, test.expectedChange, test.expectedResult,
				breakdown, change, result)
		}
	}
}
