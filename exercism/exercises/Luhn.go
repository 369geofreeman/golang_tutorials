package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

// func to reverse string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Validate a variety of identification numbers
func Valid(id string) bool {

	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	total := 0
	var digitCheck = regexp.MustCompile(`^[0-9]+$`)

	if digitCheck.MatchString(id) {
		id := Reverse(id)
		for idx := 0; idx < len(id); idx++ {
			i, _ := strconv.Atoi(string(id[idx]))
			if idx%2 != 0 {
				charInt := i * 2
				if charInt > 9 {
					charInt -= 9
				}
				total += charInt
			} else {
				total += i
			}
		}
	} else {
		return false
	}

	if total%10 == 0 {
		return true
	}

	return false
}
