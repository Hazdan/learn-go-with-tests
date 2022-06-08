package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Manuel", "")
		want := "Hello, Manuel"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Manuel", "ES")
		want := "Hola, Manuel"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}
