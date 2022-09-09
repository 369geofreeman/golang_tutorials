package raindrops

import "fmt"

// A string that contains raindrop sounds
// corresponding to certain potential factors
func Convert(number int) string {
	ret := ""
	if number%3 == 0 {
		ret += "Pling"
	}
	if number%5 == 0 {
		ret += "Plang"
	}
	if number%7 == 0 {
		ret += "Plong"
	}
	if ret == "" {
		return fmt.Sprintf("%d", number)
	}
	return ret
}
