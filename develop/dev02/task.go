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

func unpackString(s string) (string, error) {
	res := strings.Builder{}
	count := 0
	isEscape := false

	for i, char := range s {
		if isEscape {
			res.WriteRune(char)
			isEscape = false
		} else if unicode.IsDigit(char) {
			count++

			if i > 0 {
				res.WriteString(strings.Repeat(string(s[i-1]), int(char-'0')-1))
			} else {
				res.WriteRune(char)
			}
		} else if char == '\\' {
			isEscape = true
		} else {
			res.WriteRune(char)
		}
	}

	if s == "" {
		return "", nil
	} else if len(s) == count {
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
