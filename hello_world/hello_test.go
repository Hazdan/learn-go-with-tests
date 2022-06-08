package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Manuel")
	want := "Hello, Manuel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
