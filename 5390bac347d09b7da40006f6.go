package kata

import (
	"strings"
	"unicode"
)

func ToJadenCase(str string) string {
	words := strings.Fields(str)
	if len(words) == 0 {
		return ""
	}

	var res strings.Builder

	for i, word := range words {
		r := []rune(word)
		r[0] = unicode.ToUpper(r[0])

		res.WriteString(string(r))
		if i < len(words)-1 {
			res.WriteString(" ")
		}
	}

	return res.String()
}

// Or simply use: strings.Title(str),
// Not Fun Tho
