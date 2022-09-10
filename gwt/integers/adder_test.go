package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Exected '%d' but got '%d'", expected, sum)
	}
}

// Runs an example of the function | terminal: go test -v
// This will break if the code changes to make the exmaple no longer valid
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
