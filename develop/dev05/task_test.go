package main

import (
	"reflect"
	"testing"
)

func TestGrep(t *testing.T) {
	lines := []string{
		"line one",
		"line two",
		"line three",
		"line four",
		"line five",
		"line six",
		"line seven",
		"line eight",
		"line nine",
		"line ten",
	}

	tests := []struct {
		name     string
		pattern  string
		options  Options
		expected []string
	}{
		{
			name:    "simple match",
			pattern: "line three",
			options: Options{},
			expected: []string{
				"line three",
			},
		},
		{
			name:    "ignore case",
			pattern: "LINE THREE",
			options: Options{ignoreCase: true},
			expected: []string{
				"line three",
			},
		},
		{
			name:    "invert match",
			pattern: "line three",
			options: Options{invert: true},
			expected: []string{
				"line one",
				"line two",
				"line four",
				"line five",
				"line six",
				"line seven",
				"line eight",
				"line nine",
				"line ten",
			},
		},
		{
			name:     "fixed match",
			pattern:  "line",
			options:  Options{fixed: true},
			expected: []string{},
		},
		{
			name:    "line number",
			pattern: "line three",
			options: Options{lineNum: true},
			expected: []string{
				"3:line three",
			},
		},
		{
			name:    "count",
			pattern: "line",
			options: Options{count: true},
			expected: []string{
				"10",
			},
		},
		{
			name:    "context",
			pattern: "line five",
			options: Options{context: 1},
			expected: []string{
				"line four",
				"line five",
				"line six",
			},
		},
		{
			name:    "after",
			pattern: "line five",
			options: Options{after: 2},
			expected: []string{
				"line five",
				"line six",
				"line seven",
			},
		},
		{
			name:    "before",
			pattern: "line five",
			options: Options{before: 2},
			expected: []string{
				"line three",
				"line four",
				"line five",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := grep(tt.pattern, lines, tt.options)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
