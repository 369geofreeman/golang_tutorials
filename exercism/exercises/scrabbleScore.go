package scrabble

import "strings"

// Given a word, 'Score' computes the Scrabble score for that word.
func Score(word string) int {
	scoreMap := map[string]int{"d": 2, "g": 2, "b": 3, "c": 3, "m": 3, "p": 3, "f": 4,
		"h": 4, "v": 4, "w": 4, "y": 4, "k": 5, "j": 8, "x": 8,
		"q": 10, "z": 10}
	score := 0

	for _, val := range word {
		val := strings.ToLower(string(val))
		if scoreMap[val] > 0 {
			score += scoreMap[val]
		} else {
			score += 1
		}
	}

	return score
}
