package day01

import (
	"testing"
)

//+1, +1, +1 results in  3
//+1, +1, -2 results in  0
//-1, -2, -3 results in -6

func Test_ReadFrequencies(t *testing.T) {
	tests := []struct {
		name           string
		inputs         []string
		expectedResult int
	}{
		{
			name:           "+1 +1 +1 = 3",
			inputs:         []string{"+1", "+1", "+1"},
			expectedResult: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := readFrequencies(tt.inputs)
			if result != tt.expectedResult {
				t.Errorf("got %v, want %v", result, tt.expectedResult)
			}
		})
	}
}
