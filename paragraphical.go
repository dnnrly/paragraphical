package paragraphical

import (
	"strings"
)

func getIndent(parts []string) string {
	indent := ""

	for _, v := range parts {
		if v != "" {
			break
		}
		indent += " "
	}

	return indent
}

func Format(limit int, text string) string {
	result := []string{}
	allLines := strings.Split(text, "\n")
	for _, line := range allLines {
		parts := strings.Split(line, " ")
		indent := getIndent(parts)

		if len(parts) == len(indent) {
			result = append(result, "")
			continue
		}

		if indent != "" {
			indent += " "
		}

		line = parts[0]

		for _, word := range parts[1:] {
			possible := line + " " + word
			switch {
			case len(possible) > limit:
				result = append(result, line)
				line = indent + word
			default:
				line += " " + word
			}
		}
		result = append(result, line)
	}

	output := strings.Join(result, "\n")

	return output
}
