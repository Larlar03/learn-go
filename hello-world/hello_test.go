package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("print 'Hello, name'", func(t *testing.T) {
		got := Hello("Jake Sully", "")
		want := "Hello, Jake Sully"
		assertCorrectMessage(t, got, want)
	})
	t.Run("print 'Hello, Pandora' when no name is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Pandora"
		assertCorrectMessage(t, got, want)
	})

	t.Run("print 'Hello, Pandora' in given language", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, Pandora"
		assertCorrectMessage(t, got, want)
	})

		t.Run("print 'Hello, name' in given language", func(t *testing.T) {
		got := Hello("Jake Sully", "French")
		want := "Bonjour, Jake Sully"
		assertCorrectMessage(t, got, want)
	})

	t.Run("print 'Hello, Pandora' in given language", func(t *testing.T) {
		got := Hello("", "Portuguese")
		want := "Olá, Pandora"
		assertCorrectMessage(t, got, want)
	})

		t.Run("print 'Hello, name' in given language", func(t *testing.T) {
		got := Hello("Jake Sully", "Portuguese")
		want := "Olá, Jake Sully"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}