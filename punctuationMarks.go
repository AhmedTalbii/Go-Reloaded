package goreloded

import "strings"

func PunctuationMarksLine(str string) string {
	arr := []rune(str)
	// befor
	for i := 1; i < len(arr); i++ {
		if IsPunctuation(arr[i]) {
			if i-1 > 0 && string(arr[i-1]) == " " {
				arr = DeleteElement(arr, i-1)
				i-=2
			}
		} 
	}
	
	// after 
	for i := 0; i < len(arr)-1; i++ {
		// delet
		if IsPunctuation(arr[i]) && arr[i+1] == ' ' {
			for i+2 < len(arr) && arr[i+2] == ' ' {
				arr = DeleteElement(arr,i+1)
			}
		}
		// add
		if IsPunctuation(arr[i]) && !IsPunctuation(arr[i+1]) && arr[i+1] != ' ' {
			arr = AddElement(arr, i)
		}
	}

	return string(arr)
}

func PunctuationMarks(str string) string {
	lines := strings.Split(str, "\n")
	for i := range lines {
		lines[i] = PunctuationMarksLine(lines[i])
	}
	return strings.Join(lines, "\n")
}