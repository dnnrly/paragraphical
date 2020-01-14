package paragraphical

import (
	"strings"
)

func Format(text string) string {
	parts := strings.Split(text, " ")

	result := ""
	line := parts[0]

	for _, word := range parts[1:] {
		possible := line + " " + word
		if len(possible) > 20 {
			result += line + "\n"
			line = word
		} else {
			line += " " + word
		}
	}
	result += line

	return result
}
