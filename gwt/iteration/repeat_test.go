package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

/*
Benchmarking
Run: go test -bench=.

Runs the code N times (the function decides)
Returns
---
goos: darwin
goarch: amd64
pkg: github.com/quii/learn-go-with-tests/for/v4
10000000           136 ns/op
PASS
---

What 136 ns/op means is our function takes on average 136 nanoseconds
to run. Which is pretty ok! To test this it ran it 10000000 times.
*/

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 6)
	fmt.Println(repeat)
}
