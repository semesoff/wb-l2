package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unpackString1(s string) (string, error) {
	prevChar := '\000'
	res := strings.Builder{}

	for _, char := range s {
		if unicode.IsDigit(char) {
			if prevChar == '\000' {
				return "", fmt.Errorf("invalid string")
			} else {
				res.WriteString(strings.Repeat(string(prevChar), int(char-'0')))
				prevChar = '\000'
			}
		} else {
			if prevChar == '\000' {
				prevChar = char
			} else {
				res.WriteRune(prevChar)
				prevChar = char
			}
		}
	}

	if prevChar != '\000' {
		res.WriteRune(prevChar)
	}

	return res.String(), nil
}

func unpackString(s string) (string, error) {
	prevChar := '\000'
	res := strings.Builder{}
	isEscape := false
	count := 0

	for i, char := range s {
		if isEscape {
			res.WriteRune(char)
			prevChar = char
			isEscape = false
		} else if unicode.IsDigit(char) {
			num := int(char - '0')
			count++

			if i > 0 {
				if prevChar == '\000' {
					return "", fmt.Errorf("invalid string")
				} else {
					res.WriteString(strings.Repeat(string(prevChar), num-1))
					prevChar = '\000'
				}
			} else {
				res.WriteRune(char)
			}

		} else if char == '\\' {
			isEscape = true
		} else {
			prevChar = char
			res.WriteRune(char)
		}
	}

	if s == "" {
		return "", nil
	}
	if res.Len() == count {
		return "", fmt.Errorf("invalid string")
	}

	return res.String(), nil
}

func main() {
	fmt.Println(unpackString("a4bc2d5e"))
	fmt.Println(unpackString("abcd"))
	fmt.Println(unpackString("45"))
	fmt.Println(unpackString(""))
	fmt.Println(unpackString("qwe\\4\\5"))
	fmt.Println(unpackString("qwe\\45"))
	fmt.Println(unpackString("qwe\\\\5"))
}
