package edit

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func EditHexBin(text string) string {
	hextrim := regexp.MustCompile(`\b([0-9-Fa-f]+) \(hex\)`)
	text = hextrim.ReplaceAllStringFunc(text, func(ishex string) string {
		section := strings.Split(ishex, " ")
		hex, _ := strconv.ParseInt(section[0], 16, 64)
		return fmt.Sprintf("%d", hex)
	})
	bintrim := regexp.MustCompile(`(\b[01]+\b) \(bin\)`)
	text = bintrim.ReplaceAllStringFunc(text, func(isbin string) string {
		section := strings.Split(isbin, " ")
		bin, _ := strconv.ParseInt(section[0], 2, 64)
		return fmt.Sprintf("%d", bin)
	})
	return text
}
