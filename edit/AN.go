package edit

import (
	"strings"
)

// EditAN function replaces occurrences of "a" with "an" in the input text if the following word starts with a vowel or 'h'.
func EditAN(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" && isVowelH(words[i+1][0]) {
			words[i] += "n"
		}
	}
	return strings.Join(words, " ")
}

// isVowelH checks if a character is a vowel or 'h'.
func isVowelH(ch byte) bool {
	return strings.ContainsRune("aeiouhAEIOUH", rune(ch))
}
