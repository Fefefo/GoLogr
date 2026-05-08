package parser

import "strings"

func ParseLevel(line string) string {
	parts := strings.Fields(line)
	if len(parts) > 1 {
		return parts[1]
	}
	return "UNKNOWN"
}

func ParseErrorMessage(line string) string {
	parts := strings.Fields(line)
	if len(parts) > 2 {
		return strings.Join(parts[2:], " ")
	}
	return "UNKNOWN_ERROR"
}
