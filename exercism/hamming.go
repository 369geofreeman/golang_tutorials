package hamming

import "errors"

// Calculates the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("String a & b must be the same length")
	}

	count := 0

	for idx := range a {
		if a[idx] != b[idx] {
			count += 1
		}
	}
	return count, nil
}
