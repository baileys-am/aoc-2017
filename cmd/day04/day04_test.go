package main

import (
	"testing"
)

func TestValidPassphrase(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for _, tt := range tests {
		actual := ValidatePassphrase(tt.in)

		if tt.out != actual {
			t.Errorf("failed to validate passphrase")
		}
	}
}
