package goreloded

func Solve(str string) string {
	// clean
	str = CleanSpaces(str)
	// vowels
	str = VowelsHandle(str)
	// ,;
	str = PunctuationMarks(str)
	// ()
	str = OrdersHandle(str)
	// ,;
	str = PunctuationMarksAfterHandle(str)
	// ”
	str = SingleQuotes(str)
	// clean
	str = CleanSpaces(str)
	return str
}
