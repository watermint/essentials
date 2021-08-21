package ecase

import (
	"essentials/estring/etokenizer"
	"strings"
	"unicode"
)

// TokenToCamelCase changes token to camel case
func TokenToCamelCase(s string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return strings.ToUpper(s)
	default:
		runes := []rune(s)
		runes[0] = unicode.ToUpper(runes[0])
		for i := 1; i < len(runes); i++ {
			runes[i] = unicode.ToLower(runes[i])
		}
		return string(runes)
	}
}

// ToUpperCamelCase change to upper camel case like from "powered by go-lang" to "PoweredByGoLang".
func ToUpperCamelCase(s string) string {
	input := etokenizer.AlphaNumCaseTokenizer().Tokens(s)
	if len(input) < 1 {
		return ""
	}
	output := make([]string, len(input))

	for i := 0; i < len(output); i++ {
		output[i] = TokenToCamelCase(input[i])
	}
	return strings.Join(output, "")
}

func ToLowerCamelCase(s string) string {
	input := etokenizer.AlphaNumCaseTokenizer().Tokens(s)
	if len(input) < 1 {
		return ""
	}
	output := make([]string, len(input))
	output[0] = strings.ToLower(input[0])
	if len(input) < 2 {
		return output[0]
	}
	for i := 1; i < len(output); i++ {
		output[i] = TokenToCamelCase(input[i])
	}
	return strings.Join(output, "")
}
