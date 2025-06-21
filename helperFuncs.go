package goreloded

import (
	"strings"
	"unicode"
)

func Atoi(s string) int {
	start := 0
	result := 0
	negative := false
	if s[0] == '+' || s[0] == '-' {
		start = 1
		if s[0] == '-' {
			negative = true
		}
	}
	for i := start ; i < len(s) ; i++ {
		if s[i] <'0' || s[i] >'9' {
			return 0
		}
		digit := int(s[i] - '0')
		result = result*10 + digit
	}
	if negative {
		return -result
	}

	return result
}

func CleanSpaces(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.Join(strings.Fields(line), " ")
	}
	return strings.Join(lines, "\n")
}

func DeleteElement(s []rune, i int) []rune {
	if i < 0 || i >= len(s) {
		return s
	}

	return append(s[:i], s[i+1:]...)
}

func DeletString(arr []string, i int) []string {
	if i < 0 || i >= len(arr) {
		return arr
	}

	return append(arr[:i], arr[i+1:]...)
}

func Capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	runes := []rune(str)
	start := false
	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) && !start {
			runes[i] = unicode.ToUpper(runes[i])
			start = true
		} else {
			runes[i] = unicode.ToLower(runes[i])
		}
	}
	return string(runes)
}

func AddElement(s []rune,i int) []rune {
	if i < 0 || i >= len(s) {
		return s
	}
	arr := []rune{}
	for j := 0 ; j < i+1 ; j++ {
		arr = append(arr, rune(s[j]))
	}
	arr = append(arr, ' ')
	for j := i+1 ; j < len(s) ; j++ {
		arr = append(arr, rune(s[j]))
	}
	return arr
}

func IsPunctuation(r rune) bool {
	return (string(r) == "." || string(r) == "," || string(r) == "!" || string(r) == "?" || string(r) == ";" || string(r) == ":")
}

func isVowel(str string) bool {
	if str == "" {
		return false
	}
	v := rune(strings.ToLower(string(str[0]))[0])
	if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'h' {
		return true
	}
	return false
}
