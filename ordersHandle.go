package goreloded

import (
	"fmt"
	"strconv"
	"strings"
)

func checkDouble(str string) bool {
	r := []rune(str)
	if r[0] != '(' || r[len(str)-1] != ')' {
		return false
	}
	flagsVn := []string{"(hex, ", "(bin, ", "(up, ", "(low, ", "(cap, "}
	res := true
	for _, flag := range flagsVn {
		fmt.Println(string(r[0:len(flag)]))
		if string(r[0:len(flag)]) != flag {
			res = false
		}
	}
	return res
}

func applyOrders(arr []string) []string {
	flagsNn := []string{"(hex)", "(bin)", "(up)", "(low)", "(cap)"}
	flagsVn := []string{"(hex, ", "(bin, ", "(up, ", "(low, ", "(cap, "}
	for i := 0; i < len(arr); i++ {
		v := arr[i]

		// handle first 
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
				arr = DeletString(arr, i)
				i-=1
			}
		}
		// number cases

		for _, flag := range flagsVn {
			if checkDouble(v) && i != 0 {
				switch flag {
				case "(hex, ":
					r := []rune(arr[i][len(flag):len(v)]) 
					fmt.Println(string(r))
					nbr,err := strconv.Atoi(string(r))
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(nbr)
				case "(bin, ":
					r := []rune(arr[i][len(flag):len(v)])
					fmt.Println(string(r))
					nbr,err := strconv.Atoi(string(r))
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(nbr)
				case "(up, ":
					r := []rune(arr[i][len(flag):len(v)]) 
					fmt.Println(string(r))
					nbr,err := strconv.Atoi(string(r))
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(nbr)
				case "(low, ":
					r := []rune(arr[i][len(flag):len(v)]) 
					nbr,err := strconv.Atoi(string(r))
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(nbr)
				case "(cap, ":
					r := []rune(arr[i][len(flag):len(v)]) 
					nbr,err := strconv.Atoi(string(r))
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(nbr)
				}
				arr = DeletString(arr, i)
				i-=1
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
		for i := 0; i < len(line); i++ {
			char := line[i]

			if char == ' ' && start && (i+2 >= len(line) || line[i+2] != ')') {
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
