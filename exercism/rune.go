package logs

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		if fmt.Sprintf("%c", char) == "❗" {
			return "recommendation"
		}
		if fmt.Sprintf("%c", char) == "🔍" {
			return "search"
		}
		if fmt.Sprintf("%c", char) == "☀" {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	out := []rune(log)
	for idx, val := range out {
		if val == oldRune {
			out[idx] = newRune
		}
	}
	log = string(out)
	return log
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
