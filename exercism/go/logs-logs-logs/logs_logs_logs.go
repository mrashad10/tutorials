package logs

import (
	"strings"
	"unicode/utf8"
)

var applicationsMap = map[string]string{
	"❗": "recommendation",
	"🔍": "search",
	"☀": "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {
		if app, ok := applicationsMap[string(v)]; ok {
			return app
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.Replace(log, string(oldRune), string(newRune), -1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
