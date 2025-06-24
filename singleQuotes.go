package goreloded

import (
	"strings"
	"unicode"
)

func SingleQuotesInLine(str string) string {
	r := []rune(str)
	start := false
	li := 0
	for i := 0; i < len(r); i++ {
		if i-1 >= 0 && i+1 < len(r) && string(r[i]) == "'" && unicode.IsLetter(r[i-1]) && unicode.IsLetter(r[i+1]) {
			continue
		}
		if string(r[i]) == "'" && !start {
			start = true
			li = i
		} else if string(r[i]) == "'" && start {
			if r[li+1] == ' ' {
				r = DeleteElement(r, li+1)
				i = i - 2
			} else if r[i-1] == ' ' {
				r = DeleteElement(r, i-1)
				i = i - 2
			} else {
				start = false
				continue
			}
		}
	}
	return string(r)
}

func SingleQuotes(str string) string {
	lines := strings.Split(str, "\n")
	for i := range lines {
		lines[i] = SingleQuotesInLine(lines[i])
	}
	return strings.Join(lines, "\n")
}
