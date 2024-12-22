package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
        {"abcd", "abcd", false},
        {"45", "", true},
        {"", "", false},
        {"qwe\\4\\5", "qwe45", false},
        {"qwe\\45", "qwe44444", false},
        {"qwe\\\\5", "qwe\\\\\\\\\\", false},
	}

	for _, test := range tests {
		result, err := unpackString(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("unpackString(%q) error = %v, wantErr %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("unpackString(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}