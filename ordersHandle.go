package goreloded

import (
	"strconv"
	"strings"
)

func applyOrders(arr []string) []string {
	flagsNn := []string{"(hex)", "(bin)", "(up)", "(low)", "(cap)"}
	flagsVn := []string{"(up, ", "(low, ", "(cap, "}
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
					arr[i-1] = strconv.Itoa(HexBinToDec("0123456789abcdef", arr[i-1]))
				case "(bin)":
					arr[i-1] = strconv.Itoa(HexBinToDec("01", arr[i-1]))
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
					if err != nil || nbr < 0 || len(r) == 0 {
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
					if err != nil || nbr < 0 || len(r) == 0 {
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
					if err != nil || nbr < 0 || len(r) == 0 {
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
		finalLines = append(finalLines, strings.Join(words, " "))
	}

	return strings.Join(finalLines, "\n")
}
