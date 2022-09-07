package main

import (
	"testing"
)

func TestConvertTime(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		t        string
		expected string
	}{
		{"07:05:45PM", "19:05:45"},
		{"07:05:45PM", "19:05:45"},
		{"07:05:45PM", "19:05:45"},
		{"07:05:45PM", "19:05:45"},
		{"07:05:45PM", "19:05:45"},
		{"07:05:45PM", "19:05:45"},
	}

	for _, testCase := range testTable {
		t.Run(testCase.t, func(t *testing.T) {
			t.Parallel()
			result := ConvertTime()
			if result != testCase.expected {
				t.Errorf("Wrong result.Expected %q, got %q", testCase.expected, result)
			}
		})

	}
}
