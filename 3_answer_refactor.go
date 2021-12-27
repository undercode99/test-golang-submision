func findFirstStringInBracket(str string) string {
	if len(str) == 0 {
		return ""
	}

	// find first opening bracket '('
	strRunes := []rune(str)
	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}

	// find first closing bracket ')'
	wordsAfterFirstBracket := string(strRunes[indexFirstBracketFound:len(str)])
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound < 0 {
		return ""
	}

	// return string between first bracket and closing bracket
	wordsInsideBracketsRunes := []rune(wordsAfterFirstBracket)
	wordsInsideBrackets := string(wordsInsideBracketsRunes[1:indexClosingBracketFound])

	return wordsInsideBrackets
}
