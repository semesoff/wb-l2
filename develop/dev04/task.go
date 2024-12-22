package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortString(s string) string {
	r := []rune(s)
	// sort runes
	sort.SliceStable(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string) // map of anagrams
	seen := make(map[string]bool)         // map of seen words

	// sort words and group them by sorted word
	for _, word := range words {
		// convert word to lower case
		lowerWord := strings.ToLower(word)
		// sort word
		sortedWord := sortString(lowerWord)

		// skip if word is already seen
		if _, ok := seen[lowerWord]; ok {
			continue
		}

		// if sorted word is not in anagrams map, add it
		if _, ok := anagrams[sortedWord]; !ok {
			anagrams[sortedWord] = []string{}
		}
		// add word to anagrams map
		anagrams[sortedWord] = append(anagrams[sortedWord], lowerWord)
		seen[lowerWord] = true
	}

	// remove single word anagrams
	result := make(map[string][]string)
	for _, group := range anagrams {
		// if group has more than 1 word
		if len(group) > 1 {
			sort.Strings(group)
			result[group[0]] = group
		}
	}
	return result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := findAnagrams(words)
	for key, group := range anagrams {
		fmt.Printf("%s: %v\n", key, group)
	}
}
