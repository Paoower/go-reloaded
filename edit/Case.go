package edit

import (
	"regexp"
	"strconv"
	"strings"
)

func EditCase(text string) string {
	upTrim := regexp.MustCompile(`(\b\w+\b) \(up\)`)
	text = upTrim.ReplaceAllStringFunc(text, func(match string) string {
		section := strings.Split(match, " ")
		return strings.ToUpper(section[0])
	})

	lowTrim := regexp.MustCompile(`(\b\w+\b) \(low\)`)
	text = lowTrim.ReplaceAllStringFunc(text, func(match string) string {
		section := strings.Split(match, " ")
		return strings.ToLower(section[0])
	})

	capTrim := regexp.MustCompile(`(\b\w+\b) \(cap\)`)
	text = capTrim.ReplaceAllStringFunc(text, func(match string) string {
		section := strings.Split(match, " ")
		return strings.Title(section[0])
	})

	upNTrim := regexp.MustCompile(`\b(\w+\b\s+){0,}\b\w+\b\s+\(up, \d+\)`)
	lowNTrim := regexp.MustCompile(`\b(\w+\b\s+){0,}\b\w+\b\s+\(low, \d+\)`)
	capNTrim := regexp.MustCompile(`\b(\w+\b\s+){0,}\b\w+\b\s+\(cap, \d+\)`)

	text = processMultipleCase(text, upNTrim, strings.ToUpper)
	text = processMultipleCase(text, lowNTrim, strings.ToLower)
	text = processMultipleCase(text, capNTrim, strings.Title)

	return text
}

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
