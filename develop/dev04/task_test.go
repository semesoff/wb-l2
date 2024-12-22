package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		words    []string
		expected map[string][]string
	}{
		{
			[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "слиток"},
			map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "слиток", "Слиток"},
			map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "слиток", "Слиток", "СЛИТОК"},
			map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "слиток", "Слиток", "СЛИТОК", "СЛИТОК"},
			map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	k := 1
	for _, test := range tests {
		t.Run(fmt.Sprintf("Test: %d", k), func(t *testing.T) {
			result := findAnagrams(test.words)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("findAnagrams(%v) = %v, want %v", test.words, result, test.expected)
			}
		})
		k++
	}
}
