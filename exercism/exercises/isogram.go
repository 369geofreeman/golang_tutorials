package isogram

import "strings"

func IsIsogram(word string) bool {
	s := ""
	for _, char := range word {
		if string(char) == "-" || string(char) == " " {
			continue
		}
		for _, charTwo := range s {
			c1, c2 := strings.ToLower(string(char)), strings.ToLower(string(charTwo))
			if c1 == c2 {
				return false
			}
		}
		s += string(char)
	}
	return true
}
