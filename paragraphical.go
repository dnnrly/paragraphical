package paragraphical

import (
	"strings"
)

func Format(limit int, text string) string {
	result := []string{}
	allLines := strings.Split(text, "\n")
	for _, line := range allLines {
		parts := strings.Split(line, " ")

		line = parts[0]

		for _, word := range parts[1:] {
			possible := line + " " + word
			switch {
			case len(possible) > limit:
				result = append(result, line)
				line = word
			default:
				line += " " + word
			}
		}
		result = append(result, line)
	}

	output := strings.Join(result, "\n")

	return output
}
