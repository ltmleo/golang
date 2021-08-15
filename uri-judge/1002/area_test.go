package main

import (
	"testing"
)

func TestArea(t *testing.T) {
    t.Run("Case 1", func(t *testing.T) {
        assertArea(t, 2.00, 12.5664)
    })
	t.Run("Case 2", func(t *testing.T) {
        assertArea(t, 100.64, 31819.3103)
    })
	t.Run("Case 3", func(t *testing.T) {
        assertArea(t, 150.00, 70685.7750 )
    })
}
func TestRound(t *testing.T) {
    t.Run("Case 1", func(t *testing.T) {
        assertRound(t, 12.56636, 10000, 12.5664)
    })
	t.Run("Case 2", func(t *testing.T) {
        assertRound(t, 12.56634, 10000, 12.5663)
    })
	t.Run("Case 3", func(t *testing.T) {
        assertRound(t, 12.15, 10, 12.2 )
    })
	t.Run("Case 4", func(t *testing.T) {
        assertRound(t, 12.1551, 10, 12.2 )
    })
}

func assertRound(t *testing.T, number float64, precision float64, expectRounded float64) {
	t.Helper()
    rounded := Round(number, precision)

    if expectRounded != rounded {
        t.Errorf("got %g want %g", rounded, expectRounded)
    }
}

func assertArea(t *testing.T, radius float64, expectedArea float64) {
	t.Helper()
    area := Area(radius)

    if expectedArea != area {
        t.Errorf("got %g want %g", area, expectedArea)
    }
}