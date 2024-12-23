package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Options store flags values
type Options struct {
	fields       string
	delimiter    string
	separated    bool
	filedIndexes []int
}

// parseFlags parses command line flags and returns Options
func parseFlags() (Options, error) {
	var options Options
	// set flags
	flag.StringVar(&options.fields, "f", "", "fields")
	flag.StringVar(&options.delimiter, "d", "\t", "delimiter")
	flag.BoolVar(&options.separated, "s", false, "separated")
	// parse flags
	flag.Parse()
	// parse fields into fieldIndexes
	var err error
	options.filedIndexes, err = parseFields(options.fields)
	return options, err
}

// readStdin reads lines from stdin and returns them as slice of strings
func readStdin() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	// iter over lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// выбираем столбцы в строках по индексам
func selectColumns(columns []string, indexes []int) []string {
	// если оден столбец - возвращаем его
	if len(columns) == 1 {
		return columns
	}
	var selected []string
	// проходимся по индексам и проверяем, что индекс меньше длины строки
	for _, index := range indexes {
		if index < len(columns) {
			selected = append(selected, columns[index])
		}
	}
	return selected
}

// основная функция обработки строк
func cut(lines []string, options Options) []string {
	var result []string
	for _, line := range lines {
		// если флаг -s, то проверяем наличие разделителя, если его нет, то пропускаем строку
		if options.separated && !strings.Contains(line, options.delimiter) {
			continue
		}
		// разбиваем строку на столбцы по options.delimiter
		columns := strings.Split(line, options.delimiter)
		// выбираем нужные столбцы по индексам
		selectedColumns := selectColumns(columns, options.filedIndexes)
		// добавляем в результат строку из выбранных столбцов
		result = append(result, strings.Join(selectedColumns, options.delimiter))
	}
	return result
}

// parseFields
func parseFields(fields string) ([]int, error) {
	var indexes []int
	for _, field := range strings.Split(fields, ",") {
		var index int
		fmt.Sscanf(field, "%d", &index)
		if index == 0 {
			return nil, fmt.Errorf("invalid field index: %s", field)
		}
		// индексы начинаются с 1
		index--
		indexes = append(indexes, index)
	}
	return indexes, nil
}

func main() {
	// parse flags
	options, err := parseFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// read input from stdin
	lines, err := readStdin()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// process
	fmt.Println(lines, options)
	result := cut(lines, options)

	// print result
	for _, line := range result {
		fmt.Println(line)
	}
}
