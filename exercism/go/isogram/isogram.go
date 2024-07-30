package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks if a word is an isogram.
//
// It takes a string as a parameter representing the word to be checked.
// It returns a boolean indicating whether the word is an isogram.
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, char := range strings.ToLower(word) {
		if char == ' ' || char == '-' {
			continue
		}
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

type blank struct{}

// IsIsogram accepts a phrase and returns if it is an isogram.
func IsIsogram2(phrase string) bool {
	lookup := make(map[rune]blank)
	for _, chr := range phrase {
		if !unicode.IsLetter(chr) {
			continue
		}
		ltr := unicode.ToLower(chr)
		if _, exists := lookup[ltr]; exists {
			return false
		}
		lookup[ltr] = blank{}
	}
	return true
}

func IsIsogram3(phrase string) bool {
	phrase = strings.ToLower(phrase)
	for _, chr := range phrase {
		if !unicode.IsLetter(chr) {
			continue
		}
		if strings.Count(phrase, string(chr)) > 1 {
			return false
		}
	}
	return true
}

func IsIsogram4(phrase string) bool {
	theEnd := len(phrase)
	var letterFlags uint32 = 0

	for i := 0; i < theEnd; i++ {
		letter := phrase[i]
		if letter > 96 && letter < 123 {
			if (letterFlags & (1 << (letter - 'a'))) != 0 {
				return false
			} else {
				letterFlags |= (1 << (letter - 'a'))
			}
		} else if letter > 64 && letter < 91 {
			if (letterFlags & (1 << (letter - 'A'))) != 0 {
				return false
			} else {
				letterFlags |= (1 << (letter - 'A'))
			}
		}
	}
	return true
}
