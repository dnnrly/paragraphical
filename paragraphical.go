package paragraphical

import (
	"strings"
)

func Format(limit int, text string) string {
	parts := strings.Split(text, " ")

	result := ""
	line := parts[0]

	for _, word := range parts[1:] {
		possible := line + " " + word
		switch {
		case len(possible) > limit:
			result += line + "\n"
			line = word
			//		case len(possible) == limit:
			//			result += possible + "\n"
			//			line = word
		default:
			line += " " + word
		}
	}
	result += line

	return result
}
