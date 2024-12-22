package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Options struct to store flags values
type Options struct {
	column       int
	numeric      bool
	reverse      bool
	unique       bool
	month        bool
	ignoreSpaces bool
	checkSorted  bool
	humanNumeric bool
}

// parse function to parse flags and file name
func parse() (Options, string, error) {
	options := parseFlags()
	fileName, err := parseFileName()
	if err != nil {
		return Options{}, "", err
	}
	return options, fileName, nil
}

// parseFlags function to parse flags
func parseFlags() Options {
	var options Options
	// parse flags and store values in options
	flag.IntVar(&options.column, "k", 0, "column to sort by")
	flag.BoolVar(&options.numeric, "n", false, "sort by numeric value")
	flag.BoolVar(&options.reverse, "r", false, "sort in reverse order")
	flag.BoolVar(&options.unique, "u", false, "do not output duplicate lines")
	flag.BoolVar(&options.month, "M", false, "sort by month name")
	flag.BoolVar(&options.ignoreSpaces, "b", false, "ignore trailing spaces")
	flag.BoolVar(&options.checkSorted, "c", false, "check if data is sorted")
	flag.BoolVar(&options.humanNumeric, "h", false, "sort by numeric value with suffixes")
	flag.Parse()
	return options
}

// parseFileName function to parse file name from arguments
func parseFileName() (string, error) {
	// read file name from arguments
	args := flag.Args()
	// if count of arguments is more than 0, return first argument
	if len(args) > 0 {
		return args[0], nil
	}
	// if no arguments provided, return error
	return "", fmt.Errorf("no file name provided")
}

// read function to read input from file or stdin
func readInput(fileName string) ([]string, error) {
	var lines []string         // store lines
	var scanner *bufio.Scanner // scanner to read from file or stdin

	// if fileName is not empty, open file and read from it
	if fileName != "" {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		// close file after reading
		defer file.Close()
		// create scanner to read from file
		scanner = bufio.NewScanner(file)
	} else {
		// create scanner to read from stdin
		scanner = bufio.NewScanner(os.Stdin)
	}

	// read lines from scanner
	for scanner.Scan() {
		// append line to lines
		lines = append(lines, scanner.Text())
	}

	// check if there was an error while reading
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// checkSorted function to check if lines are sorted
func checkSorted(lines []string, options Options) bool {
	// iterate over lines
	for i := 1; i < len(lines); i++ {
		// if reverse is true, check if previous line is less than current line
		if options.reverse {
			if lines[i-1] < lines[i] {
				return false
			}
		} else {
			if lines[i-1] > lines[i] {
				return false
			}
		}
	}
	return true
}

// uniqueLines function to remove duplicate lines
func uniqueLines(lines []string) []string {
	// store unique lines
	hash := make(map[string]bool)
	// store result
	var result []string
	// iter over lines
	for _, line := range lines {
		// if line isn't in hash, add it to result
		if _, ok := hash[line]; !ok {
			hash[line] = true
			result = append(result, line)
		}
	}
	return result
}

// parseHumanNumeric function to parse human numeric values
func parseHumanNumeric(s string) (float64, error) {
	// map to store multipliers
	multipliers := map[byte]float64{
		'K': 1e3,
		'M': 1e6,
		'G': 1e9,
		'T': 1e12,
		'P': 1e15,
		'E': 1e18,
		'Z': 1e21,
		'Y': 1e24,
	}

	if len(s) == 0 {
		return 0, fmt.Errorf("empty string")
	}

	// check if last character is in multipliers
	lastChar := s[len(s)-1]
	if multiplier, ok := multipliers[lastChar]; ok {
		// parse number and multiply it by multiplier
		num, err := strconv.ParseFloat(s[:len(s)-1], 64)
		if err != nil {
			return 0, err
		}
		return num * multiplier, nil
	}
	return strconv.ParseFloat(s, 64)
}

// sortLines function to sort lines based on flags
func sortLines(lines []string, options Options) []string {
	// -b
	if options.ignoreSpaces {
		for i, line := range lines {
			lines[i] = strings.TrimRight(line, " ")
		}
	}

	// -u
	if options.unique {
		lines = uniqueLines(lines)
	}

	// sort lines
	sort.SliceStable(lines, func(i, j int) bool {
		a, b := lines[i], lines[j]

		// -k
		if options.column > 0 {
			// split lines by spaces
			aFields := strings.Fields(a)
			bFields := strings.Fields(b)
			// check if column is less than length of fields
			if options.column-1 < len(aFields) && options.column-1 < len(bFields) {
				a, b = aFields[options.column-1], bFields[options.column-1]
			}
		}

		// -n
		if options.numeric {
			// parse a, b to float64
			aNum, aErr := strconv.ParseFloat(a, 64)
			bNum, bErr := strconv.ParseFloat(b, 64)
			if aErr == nil && bErr == nil {
				// -r
				if options.reverse {
					return aNum > bNum
				}
				return aNum < bNum
			}
		}

		// -h
		if options.humanNumeric {
			aNum, aErr := parseHumanNumeric(a)
			bNum, bErr := parseHumanNumeric(b)
			if aErr == nil && bErr == nil {
				if options.reverse {
					return aNum > bNum
				}
				return aNum < bNum
			}
		}

		// -M
		if options.month {
			// parse a, b to time.Time
			aTime, aErr := time.Parse("Jan", a)
			bTime, bErr := time.Parse("Jan", b)
			if aErr == nil && bErr == nil {
				// -r
				if options.reverse {
					return aTime.After(bTime)
				}
				return aTime.Before(bTime)
			}
		}

		// -r
		if options.reverse {
			return a > b
		}
		return a < b
	})

	return lines
}

func sortFunc() {
	// Parse flags
	options, fileName, err := parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read input
	lines, err := readInput(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// -c Check if input is sorted
	if options.checkSorted {
		if checkSorted(lines, options) {
			fmt.Println("The input is sorted.")
		} else {
			fmt.Println("The input is not sorted.")
		}
		return
	}

	// Sort lines
	sortedLines := sortLines(lines, options)
	for _, line := range sortedLines {
		fmt.Println(line)
	}
}

func main() {
	sortFunc()
}
