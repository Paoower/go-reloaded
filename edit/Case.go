package edit

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// EditCase function performs various case modifications to the input text based on specified patterns.
// It supports three patterns: (up) for uppercase, (low) for lowercase, and (cap) for capitalized.
func EditCase(text string) string {
	// Match uppercase, lowercase, and capitalized patterns
	upTrim := regexp.MustCompile(`(\b\w+\b) \(up\)`)
	lowTrim := regexp.MustCompile(`(\b\w+\b) \(low\)`)
	capTrim := regexp.MustCompile(`(\b\w+\b) \(cap\)`)

	// Replace uppercase patterns with uppercase text
	text = upTrim.ReplaceAllStringFunc(text, func(match string) string {
		section := strings.Split(match, " ")
		return strings.ToUpper(section[0])
	})

	// Replace lowercase patterns with lowercase text
	text = lowTrim.ReplaceAllStringFunc(text, func(match string) string {
		section := strings.Split(match, " ")
		return strings.ToLower(section[0])
	})

	// Replace capitalized patterns with capitalized text
	text = capTrim.ReplaceAllStringFunc(text, func(match string) string {
		section := strings.Split(match, " ")
		return capitalizeWords(section[0])
	})

	// Match patterns with specified number of words and case
	upNTrim := regexp.MustCompile(`\b(\w+\b\s+){0,}\b\w+\b\s+\(up, \d+\)`)
	lowNTrim := regexp.MustCompile(`\b(\w+\b\s+){0,}\b\w+\b\s+\(low, \d+\)`)
	capNTrim := regexp.MustCompile(`\b(\w+\b\s+){0,}\b\w+\b\s+\(cap, \d+\)`)

	// Process patterns with specified number of words and case
	text = processMultipleCase(text, upNTrim, strings.ToUpper)
	text = processMultipleCase(text, lowNTrim, strings.ToLower)
	text = processMultipleCase(text, capNTrim, capitalizeWords)

	return text
}

// processMultipleCase applies case modifications to patterns with specified number of words (func,nb).
func processMultipleCase(text string, Trim *regexp.Regexp, caseFunc func(string) string) string {
	matches := Trim.FindAllStringSubmatchIndex(text, -1)
	if len(matches) == 0 {
		return text
	}

	var result strings.Builder
	prevEnd := 0

	for _, match := range matches {
		start, end := match[0], match[1]
		fullMatch := text[start:end]
		numStr := regexp.MustCompile(`\d+`).FindString(fullMatch)
		num, _ := strconv.Atoi(numStr)

		result.WriteString(text[prevEnd:start])
		transformed := applyCaseToNWords(fullMatch, num, caseFunc)
		result.WriteString(transformed)
		prevEnd = end
	}
	result.WriteString(text[prevEnd:])

	return result.String()
}

// applyCaseToNWords applies case modification to a specified number of words in a segment of text.
func applyCaseToNWords(segment string, num int, caseFunc func(string) string) string {
	words := strings.Fields(segment)
	trimIndex := len(words) - 2
	words = words[:trimIndex]
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	for i := 0; i < num && i < len(words); i++ {
		words[i] = caseFunc(words[i])
	}
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

// capitalizeWords capitalizes the first letter of each word and puts the rest of the letters in lowercase.
func capitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}
