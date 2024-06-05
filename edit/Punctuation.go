package edit

import (
	"strings"
)

// EditPunctuation function manipulates the punctuation in the input text to ensure proper spacing and placement.
func EditPunctuation(text string) string {
	puncs := []string{",", ".", "!", "?", ":", ";"}
	s := strings.Fields(text)

	// Handle punctuation in the middle of a string connecting to word after
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) != punc {
				s[i-1] += punc
				s[i] = word[1:]
			}
		}
	}

	// Handle punctuation at the end of a string
	for i, word := range s {
		for _, punc := range puncs {
			if (string(word[0]) == punc) && (s[len(s)-1] == s[i]) {
				s[i-1] += word
				s = s[:len(s)-1]
			}
		}
	}

	// Handle punctuation in the middle of a string
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) == punc && s[i] != s[len(s)-1] {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return strings.Join(s, " ")
}
