package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("given a name is supplied, print 'hello name'", func(t *testing.T) {
		got := Hello("Jake Sully")
		want := "Hello, Jake Sully"
		assertCorrectMessage(t, got, want)
	})
	t.Run("given no name is supplied, print 'Hello, Pandora'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Pandora"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}