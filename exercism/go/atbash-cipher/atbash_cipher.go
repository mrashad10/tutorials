package atbash

import "strings"

func cipher(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return 'z' - (char - 'a')
	}
	if char >= 'A' && char <= 'Z' {
		return 'Z' - (char - 'A')
	}

	if char >= '0' && char <= '9' {
		return char
	}

	// Not a letter or digit return an empty string
	return ' '
}

func Atbash(s string) string {
	ciphered := strings.Map(cipher, s)
	ciphered = strings.ReplaceAll(ciphered, " ", "")
	ciphered = strings.ToLower(ciphered)

	// Add spaces
	var result string
	for i := 0; i < len(ciphered); i++ {
		if i != 0 && i%5 == 0 {
			result += string(' ')
		}

		result += string(ciphered[i])
	}

	return result
}
