package edit

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// EditHexBin function replaces hexadecimal and binary representations in the input text with their decimal equivalents if (bin) and (hex) are placed behind.
func EditHexBin(text string) string {
	// Match hexadecimal representations
	hextrim := regexp.MustCompile(`\b([0-9-Fa-f]+) \(hex\)`)
	text = hextrim.ReplaceAllStringFunc(text, func(match string) string {
		// Split the matched string into sections
		section := strings.Split(match, " ")
		// Parse hexadecimal to decimal
		hex, _ := strconv.ParseInt(section[0], 16, 64)
		return fmt.Sprintf("%d", hex)
	})
	// Match binary representations
	bintrim := regexp.MustCompile(`(\b[01]+\b) \(bin\)`)
	text = bintrim.ReplaceAllStringFunc(text, func(match string) string {
		// Split the matched string into sections
		section := strings.Split(match, " ")
		// Parse binary to decimal
		bin, _ := strconv.ParseInt(section[0], 2, 64)
		return fmt.Sprintf("%d", bin)
	})
	return text
}
