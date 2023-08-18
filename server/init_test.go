package server

import (
	"errors"
	"testing"
)

func TestParsePackSizes(t *testing.T) {
	tests := []struct {
		name          string
		packSizesStr  string
		expectedSizes []int
		expectedError error
	}{
		{
			name:          "Empty pack sizes string",
			packSizesStr:  "",
			expectedSizes: []int{5000, 2000, 1000, 500, 250},
			expectedError: nil,
		},
		{
			name:          "Valid pack sizes string",
			packSizesStr:  "300,600,1200",
			expectedSizes: []int{1200, 600, 300},
			expectedError: nil,
		},
		{
			name:          "Invalid values",
			packSizesStr:  "test",
			expectedSizes: []int{},
			expectedError: errors.New("invalid input"),
		},
		{
			name:          "With duplicates",
			packSizesStr:  "5000,5000,5000",
			expectedSizes: []int{5000},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualSizes, err := parsePackSizes(test.packSizesStr)
			if err != nil {
				if test.expectedError != nil && err.Error() != test.expectedError.Error() {
					t.Errorf("Expected error %v, but got %v", test.expectedError, err)
					return
				}
			}
			if len(actualSizes) != len(test.expectedSizes) {
				t.Errorf("Expected %d sizes, but got %d sizes", len(test.expectedSizes), len(actualSizes))
			}
			for i, size := range actualSizes {
				if size != test.expectedSizes[i] {
					t.Errorf("Expected size %d at index %d, but got %d", test.expectedSizes[i], i, size)
				}
			}
		})
	}
}
