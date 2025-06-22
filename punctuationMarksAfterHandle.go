package goreloded

func PunctuationMarksAfterHandle(str string) string {
	arr := []rune(str)
	// befor
	for i := 1; i < len(arr); i++ {
		if IsPunctuation(arr[i]) {
			if i-1 >= 0 && string(arr[i-1]) == " " && string(arr[i-2]) == ")" {
				arr = DeleteElement(arr, i-1)
				i-=2
			} else if i-1 >= 0 && string(arr[i-1]) == " " && string(arr[i-2]) != ")" {
				arr = DeleteElement(arr, i-1)
				i-=2
			}
		} 
	}
	
	// after 
	for i := 0; i < len(arr)-1; i++ {
		if IsPunctuation(arr[i]) && !IsPunctuation(arr[i+1]) && arr[i+1] != ' ' {
			arr = AddElement(arr, i)
		}
	}

	return string(arr)
}