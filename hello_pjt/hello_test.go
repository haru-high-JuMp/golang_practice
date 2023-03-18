package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Haruki")
	want := "Hello, Haruki"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}