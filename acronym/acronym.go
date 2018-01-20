package acronym

import (
	"strings"
)

// Abbreviate convert a phrase to its acronym.
func Abbreviate(phrase string) string {
	chunks := strings.FieldsFunc(phrase, Split)
	acronym := ""
	for _, word := range chunks {
		acronym = acronym + strings.ToUpper(string(word[0]))
	}
	return acronym
}

func Split(r rune) bool {
	return r == ' ' || r == '-'
}
