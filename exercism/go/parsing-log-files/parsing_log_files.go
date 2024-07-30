package parsinglogfiles

import (
	"fmt"
	"regexp"
)

// validLinesRegex is a regular expression that matches lines starting with
// [INF], [ERR], or [WRN].
var validLinesRegex = regexp.MustCompile(`^(\[INF\]|\[ERR\]|\[WRN\])`)

// lineSeparatorRegex is a regular expression that matches line separators
// enclosed in <>.
var lineSeparatorRegex = regexp.MustCompile(`<[~*=-]*>`)

// passwordRegex is a regular expression that matches case-insensitive strings
// containing the word "password".
var passwordRegex = regexp.MustCompile(`(?i)"[a-z ]*password[a-z ]*"`)

// endOfLineRegex is a regular expression that matches strings starting with
// "end-of-line" followed by one or more digits.
var endOfLineRegex = regexp.MustCompile(`end-of-line\d+`)

// tagWithUserNameRegex is a regular expression that matches strings starting
// with "User" followed by one or more word characters.
var tagWithUserNameRegex = regexp.MustCompile(`User\s+(\w+)`)

// IsValidLine checks if the given text is a valid line, starting with [INF],
// [ERR], or [WRN].
//
// text: the line of text to be validated. Returns: true if the text is valid,
// false otherwise.
func IsValidLine(text string) bool {
	return validLinesRegex.MatchString(text)
}

// SplitLogLine splits a log line into multiple strings based on the old style
// line separator.
//
// It takes a single parameter 'text' of type string, which represents the log
// line that needs to be split. It returns a slice of strings, containing the
// individual parts of the log line.
func SplitLogLine(text string) []string {
	return lineSeparatorRegex.Split(text, -1)
}

// CountQuotedPasswords counts the number of lines in the given string slice
// that contain a quoted password.
//
// Parameters: - lines: a slice of strings representing the lines to be
// checked.
//
// Returns: - an integer representing the count of lines that match the
// password regular expression.
func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if passwordRegex.MatchString(line) {
			count++
		}
	}

	return count
}

// RemoveEndOfLineText removes the "end-of-line" compined with any number text
// from the given string.
//
// text: the string to remove the end of line text from. returns: the string
// with the end of line text removed.
func RemoveEndOfLineText(text string) string {
	return endOfLineRegex.ReplaceAllString(text, "")
}

// TagWithUserName tags each line in the given slice of strings
// with the username, if it is found in the line.
//
// The function takes a parameter `lines` of type []string, which
// is a slice of strings representing the lines to be tagged.
//
// The function returns a []string, which is a new slice of strings
// with the tagged lines.
func TagWithUserName(lines []string) []string {
	taggedLines := make([]string, len(lines))
	for i, line := range lines {
		if founds := tagWithUserNameRegex.FindStringSubmatch(line); founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
		taggedLines[i] = line
	}
	return taggedLines
}
