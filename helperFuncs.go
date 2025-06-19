package goreloded

import (
	"strings"
	"strconv"
)

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

func convertToDecimal(s string, base int) (int64, error) {
    return strconv.ParseInt(s, base, 64)
}