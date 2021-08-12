package main

import (
	"testing"
	"bytes"
)

func TestAdder(t *testing.T) {
    t.Run("Case 1", func(t *testing.T) {
        assertAdder(t, 10, 9, 19)
    })
	t.Run("Case 2", func(t *testing.T) {
        assertAdder(t, -10, 4, -6)
    })
	t.Run("Case 3", func(t *testing.T) {
        assertAdder(t, 15, -7, 8)
    })
}

func TestShow(t *testing.T) {
    t.Run("Case 1", func(t *testing.T) {
        assertShow(t, 19, "X = 19\n")
    })
    t.Run("Case 2", func(t *testing.T) {
        assertShow(t, -6, "X = -6\n")
    })
    t.Run("Case 3", func(t *testing.T) {
        assertShow(t, 8, "X = 8\n")
    })
}

func assertShow(t *testing.T, result int, want string) {
	t.Helper()
	buffer := bytes.Buffer{}
    Show(&buffer, result)

    got := buffer.String()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func assertAdder(t testing.TB, x int, y int, want int) {
	t.Helper()
	got := Add(x, y)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}