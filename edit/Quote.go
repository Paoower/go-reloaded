package edit

import "regexp"

func EditQuote(text string) string {
	quotePattern := regexp.MustCompile(`\s'(\s*)([^']+?)(\s*)'\s`)
	text = quotePattern.ReplaceAllString(text, " '$2' ")
	return text
}
