package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In French", func(t *testing.T) {
		got := Hello("Monde", "French")
		want := "Bonjour, Monde"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In Japanese", func(t *testing.T) {
		got := Hello("世界", "Japanese")
		want := "こんにちは, 世界"
		assertCorrectMessage(t, got, want)
	})
}

// Testing.TB is an interface that *testing.T and *testing.B both satisfy
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() tells the suite that this method is a helper.
	// Line number is reported from our function rather than the test suite
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
