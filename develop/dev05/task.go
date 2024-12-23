package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Options store flags values
type Options struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

// parseFlags parses command line flags and returns Options
func parseFlags() Options {
	var options Options
	// set flags
	flag.IntVar(&options.after, "A", 0, "print +N lines after match")
	flag.IntVar(&options.before, "B", 0, "print +N lines before match")
	flag.IntVar(&options.context, "C", 0, "print ±N lines around match")
	flag.BoolVar(&options.count, "c", false, "print count of matching lines")
	flag.BoolVar(&options.ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&options.invert, "v", false, "invert match")
	flag.BoolVar(&options.fixed, "F", false, "fixed string match")
	flag.BoolVar(&options.lineNum, "n", false, "print line number")
	// parse flags
	flag.Parse()
	return options
}

// parse pattern & filename
func parseData() (string, string, error) {
	args := flag.Args()
	if len(args) < 2 {
		return "", "", fmt.Errorf("pattern and filename are required")
	}
	return args[0], args[1], nil
}

func grep(pattern string, lines []string, options Options) []string {
	var result []string
	var count int

	// pattern to lower case
	if options.ignoreCase {
		pattern = strings.ToLower(pattern)
	}

	// loop through lines
	for i, line := range lines {
		match := false

		// if ignore case, convert line to lower case
		if options.ignoreCase {
			line = strings.ToLower(line)
		}

		// exact match with the string
		if options.fixed {
			match = line == pattern
		} else {
			// check if line contains pattern
			match = strings.Contains(line, pattern)
		}

		// if flag -v, exclude or include line
		if options.invert {
			match = !match
		}

		if match {
			// if flag -c, increment count
			if options.count {
				count++
			} else {
				// if flag -n, print line number
				if options.lineNum {
					result = append(result, fmt.Sprintf("%d:%s", i+1, line))
				} else {
					result = append(result, line)
				}
			}

			// if flag -C, print +N lines around match
			if options.context > 0 {
				var before []string
				start := max(0, i-options.context)          // start
				end := min(i+options.context+1, len(lines)) // end
				// before i
				for j := start; j < i; j++ {
					before = append(before, lines[j])
				}
				result = append(before, result...)
				// after i
				for j := i + 1; j < end; j++ {
					result = append(result, lines[j])
				}
			}

			// if flag -A, print +N lines after match
			if options.after > 0 {
				end := min(i+options.after+1, len(lines))
				for j := i + 1; j < end; j++ {
					result = append(result, lines[j])
				}
			}

			// if flag -B, print +N lines before match
			if options.before > 0 {
				var before []string
				start := max(0, i-options.before)
				for j := start; j < i; j++ {
					before = append(before, lines[j])
				}
				result = append(before, result...)
			}
		}
	}

	if options.count {
		return []string{fmt.Sprintf("%d", count)}
	}

	return result
}

func readFile(filename string) ([]string, error) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// close file
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	// read file line by line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// check for errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func main() {
	// parse flags
	options := parseFlags()

	// parse pattern & filename
	pattern, filename, err := parseData()
	if err != nil {
		fmt.Println(err)
		return
	}

	// read file
	lines, err := readFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// grep
	res := grep(pattern, lines, options)

	// print result
	for _, line := range res {
		fmt.Println(line)
	}
}
