package goreloded

import (
	"strconv"
	"strings"
)

func applyOrders(arr []string) {
	flagsNn := []string{"(hex)", "(bin)", "(up)", "(low)", "(cap)"}
	// flagsVn := []string{"(hex, ", "(bin, ", "(up, ", "(low, ", "(cap, "}
	for i := 0; i < len(arr); i++ {
	v := arr[i]

	for _, flag := range flagsNn {
		if strings.Contains(v, flag) && i != 0 {
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
					arr[i-1] = strings.ToUpper(string(arr[i-1][0])) + strings.ToLower(arr[i-1][1:])
				}
			}
			arr[i] = strings.Replace(arr[i], flag, "", 1)
			if len(arr[i]) == 0 {
				arr = DeletString(arr, i)
				i--
			}
 		}
	}
}

}

func OrdersHandle(str string) string {
	res := []string{}
	word := ""
	start := false
	for i, char := range str {
		if (char == ' ') && start && i+2 < len(str) && str[i+2] != ')' {
			res = append(res, word)
			word = ""
			start = false
		} else if (char == ' ') && !start {
			continue
		} else {
			word += string(char)
			start = true
		}
	}
	if len(word) != 0 {
		res = append(res, word)
	}
	applyOrders(res)
	return strings.Join(res, " ")
}
