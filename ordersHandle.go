package goreloded

import (
	"fmt"
	"strconv"
	"strings"
)

func IsFlagVn(s string) (bool, string, int) {
	paramKeys := []string{"up", "low", "cap"}

	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		content := strings.TrimSuffix(strings.TrimPrefix(s, "("), ")")
		parts := strings.SplitN(content, ", ", 2)

		if len(parts) == 2 {
			flagName := strings.TrimSpace(parts[0])
			numberStr := strings.TrimSpace(parts[1])

			for _, key := range paramKeys {
				if flagName == key {
					if nbr, err := strconv.Atoi(numberStr); err == nil {
						return true, key, nbr
					} else {
						fmt.Println(err)
					}
				}
			}
		}
	}

	return false, "", -1
}

func IsFlagNn(str string) bool {
	return (str == "(hex)" || str == "(bin)" || str == "(up)" || str == "(low)" || str == "(cap)")
}

func applyOrders(arr []string) []string {
	// fix the flags
	for i := 0; i < len(arr)-1; i++ {
		if strings.HasPrefix(arr[i], "(") && strings.HasSuffix(arr[i], ",") && strings.HasSuffix(arr[i+1], ")") {
			arr[i] = arr[i] + " " + arr[i+1]
			arr = DeletString(arr, i+1)
		}
	}

	counter := 0
	for len(arr) > 0 && counter < len(arr) {
		iFv, _, _ := IsFlagVn(arr[counter])

		if !(IsFlagNn(arr[counter]) || iFv) && arr[counter] != "" {
			break
		}

		if arr[counter] == "" {
			counter++
			continue
		}
		arr = DeletString(arr, counter)
	}
	// apply orders
	for i := 1; i < len(arr); i++ {
		br := false
		if IsFlagNn(arr[i]) {
			target := i - 1
			for target >= 0 && strings.TrimSpace(arr[target]) == "" {
				if target-1 < 0 {
					br = true
					break
				}
				target--
			}
			if br {
				continue
			}
			switch arr[i] {
			case "(hex)":
				if arr[target] != "" {
					if nbr, err := strconv.ParseInt(arr[target], 16, 64); err == nil {
						arr[target] = strconv.Itoa(int(nbr))
					}
				}
			case "(bin)":
				if arr[target] != "" {
					if nbr, err := strconv.ParseInt(arr[target], 2, 64); err == nil {
						arr[target] = strconv.Itoa(int(nbr))
					}
				}
			case "(up)":
				arr[target] = strings.ToUpper(arr[target])
			case "(low)":
				arr[target] = strings.ToLower(arr[target])
			case "(cap)":
				if len(arr[target]) > 0 {
					arr[target] = Capitalize(arr[target])
				}
			}
			arr = DeletString(arr, i)
			i -= 1
		}
		if bl, key, nbr := IsFlagVn(arr[i]); bl {
			if nbr >= 0 {
				start := i - nbr
				if start < 0 {
					start = 0
				}
				for j := i - 1; j >= start; j-- {
					if arr[j] == "" {
						if start-1 >= 0 {
							start--
						}
					} else {
						switch key {
						case "up":
							arr[j] = strings.ToUpper(arr[j])
						case "low":
							arr[j] = strings.ToLower(arr[j])
						case "cap":
							if len(arr[j]) > 0 {
								arr[j] = Capitalize(arr[j])
							}
						}
					}
				}
			}
			arr = DeletString(arr, i)
			i--
		}
	}

	return arr
}

func OrdersHandle(str string) string {
	lines := strings.Split(str, "\n")
	finalLines := []string{}
	for _, line := range lines {
		words := strings.Split(line, " ")
		words = applyOrders(words)
		finalLines = append(finalLines, strings.Join(words, " "))
	}

	return strings.Join(finalLines, "\n")
}
