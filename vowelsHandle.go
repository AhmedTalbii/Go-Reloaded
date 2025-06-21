package goreloded

import "strings"

func VowelsHandle(str string) string {
	lines := strings.Split(str, "\n")
	for i := 0; i < len(lines); i++ {
		arr := strings.Fields(lines[i])
		for i := 0 ; i < len(arr)-1 ; i++ {
			if arr[i] == "a" || arr[i] == "A" && isVowel(arr[i+1]) {
				arr[i] = arr[i] + "n"
			}
		}
		lines[i] = strings.Join(arr, " ")
	}
	return strings.Join(lines, "\n")
}
