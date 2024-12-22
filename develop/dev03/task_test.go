package main

import (
	"strings"
	"testing"
)

func TestSortLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		options  Options
	}{
		{"BasicSort", "c\nb\na\n", "a\nb\nc\n", Options{}},
		{"NumericSort", "3\n1\n2\n", "1\n2\n3\n", Options{numeric: true}},
		{"ReverseSort", "a\nc\nb\n", "c\nb\na\n", Options{reverse: true}},
		{"UniqueSort", "a\na\nb\n", "a\nb\n", Options{unique: true}},
		{"MonthSort", "Jan\nMar\nFeb\n", "Jan\nFeb\nMar\n", Options{month: true}},
		{"IgnoreSpaces", "a  \nc  \nb  \n", "a\nb\nc\n", Options{ignoreSpaces: true}},
		{"ColumnSort", "a 2\nb 1\nc 3\n", "b 1\na 2\nc 3\n", Options{column: 2}},
		{"NumericColumnSort", "a 2\nb 10\nc 3\n", "a 2\nc 3\nb 10\n", Options{column: 2, numeric: true}},
		{"ReverseNumericSort", "3\n1\n2\n", "3\n2\n1\n", Options{numeric: true, reverse: true}},
		{"ReverseMonthSort", "Jan\nMar\nFeb\n", "Mar\nFeb\nJan\n", Options{month: true, reverse: true}},
		{"CheckSorted", "a\nb\nc\n", "The input is sorted.", Options{checkSorted: true}},
		{"CheckNotSorted", "c\nb\na\n", "The input is not sorted.", Options{checkSorted: true}},
		{"HumanNumericSort", "1K\n1M\n1G\n", "1K\n1M\n1G\n", Options{humanNumeric: true}},
		{"ReverseHumanNumericSort", "1G\n1M\n1K\n", "1G\n1M\n1K\n", Options{humanNumeric: true, reverse: true}},
		{"CombinedSort", "a 2\nb 1\nc 3\n", "c 3\na 2\nb 1\n", Options{column: 2, numeric: true, reverse: true}},
		{"EmptyInput", "", "", Options{}},
		{"SingleLine", "a\n", "a\n", Options{}},
		{"SingleLineWithSpaces", "a  \n", "a\n", Options{ignoreSpaces: true}},
		{"DuplicateLines", "a\na\nb\n", "a\nb\n", Options{unique: true}},
		{"MixedCase", "A\nb\nC\n", "A\nC\nb\n", Options{}},
		{"MixedCaseReverse", "A\nb\nC\n", "b\nC\nA\n", Options{reverse: true}},
		{"MixedCaseUnique", "A\nb\nC\nA\n", "A\nC\nb\n", Options{unique: true}},
		{"MixedCaseIgnoreSpaces", "A  \nb  \nC  \n", "A\nC\nb\n", Options{ignoreSpaces: true}},
		{"MixedCaseColumnSort", "A 2\nb 1\nC 3\n", "b 1\nA 2\nC 3\n", Options{column: 2}},
		{"MixedCaseNumericColumnSort", "A 2\nb 10\nC 3\n", "A 2\nC 3\nb 10\n", Options{column: 2, numeric: true}},
		{"MixedCaseReverseNumericSort", "3\n1\n2\n", "3\n2\n1\n", Options{numeric: true, reverse: true}},
		{"MixedCaseReverseMonthSort", "Jan\nMar\nFeb\n", "Mar\nFeb\nJan\n", Options{month: true, reverse: true}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lines := strings.Split(strings.TrimSpace(test.input), "\n")
			expected := strings.Split(strings.TrimSpace(test.expected), "\n")
			if test.options.checkSorted {
				if checkSorted(lines, test.options) {
					result := []string{"The input is sorted."}
					if !equal(result, expected) {
						t.Errorf("checkSorted(%q, %v) = %q, want %q", test.input, test.options, result, expected)
					}
				} else {
					result := []string{"The input is not sorted."}
					if !equal(result, expected) {
						t.Errorf("checkSorted(%q, %v) = %q, want %q", test.input, test.options, result, expected)
					}
				}
			} else {
				result := sortLines(lines, test.options)
				if !equal(result, expected) {
					t.Errorf("sortLines(%q, %v) = %q, want %q", test.input, test.options, result, expected)
				}
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
