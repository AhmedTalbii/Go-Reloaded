package goreloded

func indFromBase(base string, r rune) int {
	for i,char := range base {
		if char == r || char == rune(r+32) {
			return i
		}
	}
	return 0
}

func power(a,b int) int {
	res := 1
	for i := 1 ; i <= b ; i++ {
		res = res * a
	}
	return res
}

func HexBinToDec(base string, str string) int {
	isNegative := false
	start := 0
	if str[0] == '-' {
		isNegative = true
		start = 1
	}
	divider := []int{}
	for i := start ; i < len(str) ; i++ {
		divider = append(divider, indFromBase(base,rune(str[i])))
	}
	counter := 0
	res := 0
	maxInt := (1 << 63) - 1
	minInt := -(1 << 63)

	for i := len(divider)-1 ; i >= 0 ; i-- {
		add := (divider[i]*power(len(base),counter))
		if !isNegative && res > (maxInt - add) {
			return maxInt
		} else if isNegative && res > (-(minInt + add)) {
			return minInt
		}
		res = res + add
		counter++
	}
	if isNegative {
		return -res
	}
	return res
}