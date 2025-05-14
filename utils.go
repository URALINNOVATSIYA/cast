package cast

import (
	"unicode"
	"unicode/utf8"
)

// ucfirst makes the first letter of string upper case.
func ucfirst(str string) string {
	char, size := utf8.DecodeRuneInString(str)
	if char == utf8.RuneError {
		return str
	}
	upper := unicode.ToUpper(char)
	if char == upper {
		return str
	}
	return string(upper) + str[size:]
}

// lcfirst makes the first letter of string lower case.
func lcfirst(str string) string {
	char, size := utf8.DecodeRuneInString(str)
	if char == utf8.RuneError {
		return str
	}
	lower := unicode.ToLower(char)
	if char == lower {
		return str
	}
	return string(lower) + str[size:]
}
