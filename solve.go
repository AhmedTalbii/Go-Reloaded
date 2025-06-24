package goreloded

func Solve(str string) string {
	// ()
	str = OrdersHandle(str)
	// ,;
	str = PunctuationMarks(str)
	// ”
	str = SingleQuotes(str)
	// vowels
	str = VowelsHandle(str)

	return str
}
