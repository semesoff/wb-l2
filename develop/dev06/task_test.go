package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		options  Options
		expected string
	}{
		{
			name:     "select first column",
			input:    "col1\tcol2\tcol3\ncol4\tcol5\tcol6\ncol7\tcol8\tcol9",
			options:  Options{fields: "1"},
			expected: "col1\ncol4\ncol7",
		},
		{
			name:     "select first and third columns",
			input:    "col1\tcol2\tcol3\ncol4\tcol5\tcol6\ncol7\tcol8\tcol9",
			options:  Options{fields: "1,3"},
			expected: "col1\tcol3\ncol4\tcol6\ncol7\tcol9",
		},
		{
			name:     "select columns with custom delimiter",
			input:    "col1,col2,col3\ncol4,col5,col6\ncol7,col8,col9",
			options:  Options{fields: "2", delimiter: ","},
			expected: "col2\ncol5\ncol8",
		},
		{
			name:     "select columns with separated flag",
			input:    "col1\tcol2\tcol3\ncol4\tcol5\tcol6\ncol7\tcol8\tcol9",
			options:  Options{fields: "2", separated: true},
			expected: "col2\ncol5\ncol8",
		},
		{
			name:     "select second column with space delimiter",
			input:    "col1 col2 col3\ncol4 col5 col6\ncol7 col8 col9",
			options:  Options{fields: "2", delimiter: " "},
			expected: "col2\ncol5\ncol8",
		},
		{
			name:     "select non-existent column",
			input:    "col1\tcol2\tcol3\ncol4\tcol5\tcol6\ncol7\tcol8\tcol9",
			options:  Options{fields: "4"},
			expected: "\n\n",
		},
		{
			name:     "select first column with no delimiter",
			input:    "col1\ncol2\ncol3",
			options:  Options{fields: "1"},
			expected: "col1\ncol2\ncol3",
		},
		{
			name:     "select first column with custom delimiter and separated flag",
			input:    "col1,col2,col3\ncol4,col5,col6\ncol7,col8,col9",
			options:  Options{fields: "1", delimiter: ",", separated: true},
			expected: "col1\ncol4\ncol7",
		},
		{
			name:     "select multiple columns with custom delimiter",
			input:    "col1,col2,col3\ncol4,col5,col6\ncol7,col8,col9",
			options:  Options{fields: "1,3", delimiter: ","},
			expected: "col1,col3\ncol4,col6\ncol7,col9",
		},
		{
			name:     "select columns with mixed delimiters",
			input:    "col1,col2\tcol3\ncol4,col5\tcol6\ncol7,col8\tcol9",
			options:  Options{fields: "2", delimiter: ","},
			expected: "col2\tcol3\ncol5\tcol6\ncol8\tcol9",
		},
		{
			name:     "select columns with tab delimiter and separated flag",
			input:    "col1\tcol2\tcol3\ncol4\tcol5\tcol6\ncol7\tcol8\tcol9",
			options:  Options{fields: "2", delimiter: "\t", separated: true},
			expected: "col2\ncol5\ncol8",
		},
		{
			name:     "select columns with space delimiter and separated flag",
			input:    "col1 col2 col3\ncol4 col5 col6\ncol7 col8 col9",
			options:  Options{fields: "2", delimiter: " ", separated: true},
			expected: "col2\ncol5\ncol8",
		},
		{
			name:     "select columns with tab delimiter and no separated flag",
			input:    "col1\tcol2\tcol3\ncol4\tcol5\tcol6\ncol7\tcol8\tcol9",
			options:  Options{fields: "2", delimiter: "\t"},
			expected: "col2\ncol5\ncol8",
		},
		{
			name:     "select columns with space delimiter and no separated flag",
			input:    "col1 col2 col3\ncol4 col5 col6\ncol7 col8 col9",
			options:  Options{fields: "2", delimiter: " "},
			expected: "col2\ncol5\ncol8",
		},
		{
			name:     "select columns with mixed delimiters and no separated flag",
			input:    "col1,col2\tcol3\ncol4,col5\tcol6\ncol7,col8\tcol9",
			options:  Options{fields: "2", delimiter: ","},
			expected: "col2\tcol3\ncol5\tcol6\ncol8\tcol9",
		},
		{
			name:     "select columns with mixed delimiters and separated flag",
			input:    "col1,col2\tcol3\ncol4,col5\tcol6\ncol7,col8\tcol9",
			options:  Options{fields: "2", delimiter: ",", separated: true},
			expected: "col2\tcol3\ncol5\tcol6\ncol8\tcol9",
		},
		{
			name:     "select columns with custom delimiter and no separated flag",
			input:    "col1,col2,col3\ncol4,col5,col6\ncol7,col8,col9",
			options:  Options{fields: "2", delimiter: ","},
			expected: "col2\ncol5\ncol8",
		},
		{
			name:     "select columns with custom delimiter and separated flag",
			input:    "col1,col2,col3\ncol4,col5,col6\ncol7,col8,col9",
			options:  Options{fields: "2", delimiter: ",", separated: true},
			expected: "col2\ncol5\ncol8",
		},
		{
			name:     "select columns with tab delimiter and no separated flag",
			input:    "col1\tcol2\tcol3\ncol4\tcol5\tcol6\ncol7\tcol8\tcol9",
			options:  Options{fields: "2", delimiter: "\t"},
			expected: "col2\ncol5\ncol8",
		},
		{
			name:     "select columns with space delimiter and no separated flag",
			input:    "col1 col2 col3\ncol4 col5 col6\ncol7 col8 col9",
			options:  Options{fields: "2", delimiter: " "},
			expected: "col2\ncol5\ncol8",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Инициализируем fieldIndexes перед вызовом cut
			if test.options.delimiter == "" {
				test.options.delimiter = "\t"
			}
			test.options.filedIndexes, _ = parseFields(test.options.fields)
			lines := strings.Split(test.input, "\n")
			result := cut(lines, test.options)
			output := strings.Split(test.expected, "\n")
			if !reflect.DeepEqual(result, output) {
				t.Errorf("expected %v, got %v", output, result)
			}
		})
	}
}
