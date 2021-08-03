package main

import (
	"testing"
	"bytes"
)

func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Leo")

    got := buffer.String()
    want := "Hello, Leo"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}