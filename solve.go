package goreloded

func Solve(str string) string {
	// clean
	str = CleanSpaces(str)
	// ,;
	str = PunctuationMarks(str)
	// ()
	str = OrdersHandle(str)
	// ,;
	str = PunctuationMarks(str)
	// ”
	str = SingleQuotes(str)
	// clean
	str = CleanSpaces(str)

	return str
}
