package edit

import (
	"strings"
)

func EditAN(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" && isVowelH(words[i+1][0]) {
			words[i] += "n"
		}
	}
	return strings.Join(words, " ")
}

func isVowelH(ch byte) bool {
	return strings.ContainsRune("aeiouhAEIOUH", rune(ch))
}
