package iteration

// Returns a character repeated n times as a string
func Repeat(char string, n int) string {
	var ret string
	for i := 0; i < n; i++ {
		ret += char
	}
	return ret
}
