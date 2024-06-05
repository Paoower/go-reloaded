package edit

import "regexp"

// EditQuote function cleans up the quotation marks in the input text to ensure proper spacing.
func EditQuote(text string) string {
	// Define a regular expression to match patterns enclosed by ' '
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)

	// Replace matched patterns with the cleaned-up version and ensure spaces around quotes
	result := re.ReplaceAllStringFunc(text, func(match string) string {
		// Remove spaces within the quotes
		cleaned := re.ReplaceAllString(match, "'$1'")

		// Ensure there is exactly one space before the opening quote
		if len(cleaned) > 1 && cleaned[0] != ' ' {
			cleaned = " " + cleaned
		} else if len(cleaned) > 2 && cleaned[0] == ' ' && cleaned[1] == ' ' {
			cleaned = cleaned[1:]
		}

		// Ensure there is exactly one space after the closing quote
		if len(cleaned) > 2 && cleaned[len(cleaned)-1] != ' ' {
			cleaned = cleaned + " "
		} else if len(cleaned) > 3 && cleaned[len(cleaned)-1] == ' ' && cleaned[len(cleaned)-2] == ' ' {
			cleaned = cleaned[:len(cleaned)-1]
		}

		return cleaned
	})
	return result
}
