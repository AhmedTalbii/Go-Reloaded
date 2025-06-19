package goreloded

func SingleQuotes(str string) string {
	r := []rune(str)
	start := false
	li := 0
	for i := 0; i < len(r); i++ {
		if string(r[i]) == "'" && !start {
			start = true
			li = i
		} else if string(r[i]) == "'" && start {
			if r[li + 1] == ' ' {
				r = DeleteElement(r,li+1)
				i=i-2
			} else if r[i - 1] == ' ' {
				r = DeleteElement(r,i-1)
				i=i-2
			} else {
				start = false
				continue
			}
		}
	}
	return string(r)
}