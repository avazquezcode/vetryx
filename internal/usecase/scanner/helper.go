package scanner

// singleChars are characters that should be scanned individually, and without any extra logic.
var singleChars = map[rune]bool{
	'(': true,
	')': true,
	'{': true,
	'}': true,
	',': true,
	'.': true,
	'-': true,
	'+': true,
	';': true,
	'/': true,
	'%': true,
	'*': true,
	'!': true,
	'#': true,
}

// matchableChars are characters that might be scanned as "single chars" (if applicabe), or can be combined with certain successor characters
// in order to produce a "composite char". Eg: the char "=" can be matched with another "=", to form "==" (used for comparison).
var matchableChars = map[rune]bool{
	'=': true, // can be matched with "=" to form "equal"
	'<': true, // can be matched with "=" to form "lower or equal"; it can also be matched with ">" to form the "different" operator
	'>': true, // can be matched with "=" to form "greater or equal"
	'&': true, // can be matched with "&" to form AND operator
	'|': true, // can be matched with "|" to form OR operator
	':': true, // can be matched with "=" to form var short declarator
}

// ignorableChars are characters that can be ignored by the scanner.
var ignorableChars = map[rune]bool{
	' ':  true,
	'\t': true,
	'\r': true,
}

// isLetter returns true if the rune is a letter.
func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// isDigit returns true if the rune is a digit.
func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

// isAlphaNum returns true if the rune is alphanum (only letters and numbers).
func isAlphaNum(c rune) bool {
	return isLetter(c) || isDigit(c)
}
