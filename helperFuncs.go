package goreloded

import (
	"strings"
	"unicode"
)

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

func AddString(s []string, i int, v string) []string {
    res := make([]string, len(s)+1)
    j := 0
    for k := 0; k < len(res); k++ {
        if k == i {
            res[k] = v
        } else {
            res[k] = s[j]
            j++
        }
    }
    return res
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
