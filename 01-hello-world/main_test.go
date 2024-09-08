package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("hideyuki")
		want := "Hello, hideyuki"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, s"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
