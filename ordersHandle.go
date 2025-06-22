package goreloded

import (
	"strconv"
	"strings"
)

func IsFlag(s string) bool {
	baseFlags := []string{"(up)", "(low)", "(cap)", "(hex)", "(bin)"}
	paramKeys := []string{"up", "low", "cap"}

	for _, flag := range baseFlags {
		if s == flag {
			return true
		}
	}

	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		content := strings.TrimSuffix(strings.TrimPrefix(s, "("), ")")
		parts := strings.SplitN(content, ", ", 2)

		if len(parts) == 2 {
			flagName := strings.TrimSpace(parts[0])
			numberStr := strings.TrimSpace(parts[1])

			for _, key := range paramKeys {
				if flagName == key {
					if _, err := strconv.Atoi(numberStr); err == nil {
						return true
					}
				}
			}
		}
	}

	return false
}

func cleanFlags(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i])-1; j++ {
			if arr[i][j] == ')' && arr[i][j+1] == ' ' {
				str := strings.Fields(arr[i])
				arr[i] = str[0]
				arr = AddString(arr, i+1, str[1])
			}
		}
	}
	return arr
}

func applyOrders(arr []string) []string {
	flagsNn := []string{"(hex)", "(bin)", "(up)", "(low)", "(cap)"}
	flagsVn := []string{"(up, ", "(low, ", "(cap, "}
	arr = cleanFlags(arr)
	if len(cleanFlags(arr)) != len(arr) {
		arr = cleanFlags(arr)
	}
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		for i := 0; i < len(arr)-1; i++ {
			if strings.HasPrefix(arr[i], "(") && strings.HasSuffix(arr[i], ",") && strings.HasSuffix(arr[i+1], ")") {
				arr[i] = arr[i] + " " + arr[i+1]
				arr = DeletString(arr, i+1)
			}
		}
		// simple cases
		for _, flag := range flagsNn {
			if v == flag && i != 0 {
				switch flag {
				case "(hex)":
					if arr[i-1] != "" {
						if nbr, err := strconv.ParseInt(arr[i-1], 16, 64); err == nil {
							arr[i-1] = strconv.Itoa(int(nbr))
						}
					}
				case "(bin)":
					if arr[i-1] != "" {
						if nbr, err := strconv.ParseInt(arr[i-1], 2, 64); err == nil {
							arr[i-1] = strconv.Itoa(int(nbr))
						}
					}
				case "(up)":
					arr[i-1] = strings.ToUpper(arr[i-1])
				case "(low)":
					arr[i-1] = strings.ToLower(arr[i-1])
				case "(cap)":
					if len(arr[i-1]) > 0 {
						arr[i-1] = Capitalize(arr[i-1])
					}
				}

				v = strings.ReplaceAll(v, flag, "")
				arr[i] = v
				if len(arr[i]) == 0 {
					arr = DeletString(arr, i)
					i = i - 1
				}
			}
		}
		// number cases
		for _, flag := range flagsVn {
			if len(string(v)) > len(flag) && flag == string(v[0:len(flag)]) && i != 0 {
				switch flag {
				case "(up, ":
					r := []rune(arr[i][len(flag) : len(v)-1])
					nbr, err := strconv.Atoi(string(r))
					if err != nil || len(r) == 0 {
						if nbr != 9223372036854775807 {
							continue
						}
					}

					start := i - nbr
					if start < 0 {
						start = 0
					}
					for j := i - 1; j >= start; j-- {
						arr[j] = strings.ToUpper(arr[j])
					}
				case "(low, ":
					r := []rune(arr[i][len(flag) : len(v)-1])
					nbr, err := strconv.Atoi(string(r))
					if err != nil || len(r) == 0 {
						if nbr != 9223372036854775807 {
							continue
						}
					}

					start := i - nbr
					if start < 0 {
						start = 0
					}
					for j := i - 1; j >= start; j-- {
						arr[j] = strings.ToLower(arr[j])
					}
				case "(cap, ":
					r := []rune(arr[i][len(flag) : len(v)-1])
					nbr, err := strconv.Atoi(string(r))
					if err != nil || len(r) == 0 {
						if nbr != 9223372036854775807 {
							continue
						}
					}

					start := i - nbr
					if start < 0 {
						start = 0
					}
					for j := i - 1; j >= start; j-- {
						if len(arr[j]) > 0 {
							arr[j] = Capitalize(arr[j])
						}
					}
				}

				arr = DeletString(arr, i)
				i -= 1
			}
		}
	}
	return arr
}

func OrdersHandle(str string) string {
	lines := strings.Split(str, "\n")
	finalLines := []string{}
	for _, line := range lines {
		words := []string{}
		word := ""
		start := false
		runes := []rune(line)
		for i := 0; i < len(runes); i++ {
			char := runes[i]

			if char == ' ' && start && (i+2 >= len(runes) || runes[i+2] != ')') {
				words = append(words, word)
				word = ""
				start = false
			} else if char == ' ' && !start {
				continue
			} else {
				word += string(char)
				start = true
			}
		}
		if len(word) > 0 {
			words = append(words, word)
		}

		words = applyOrders(words)
		if len(words) != 0 && IsFlag(words[0]) {
			words = DeletString(words, 0)
		}

		finalLines = append(finalLines, strings.Join(words, " "))
	}

	return strings.Join(finalLines, "\n")
}
