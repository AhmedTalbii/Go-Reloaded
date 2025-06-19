package goreloded

func Solve(str string) string {
	// clean 
	str = CleanSpaces(str)
	// ''
	str = SingleQuotes(str)
	// ,;
	str = PunctuationMarks(str)
	// ()
	str = OrdersHandle(str)
	return str
}