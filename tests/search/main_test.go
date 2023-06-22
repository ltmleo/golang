package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestTestFunction(t *testing.T) {
	t.Run("Test with 'OK' result", func(t *testing.T) {
		expectedResult := "OK"
		result := test()
		if result != expectedResult {
			t.Errorf("Expected '%s', but got '%s'", expectedResult, result)
		}
	})

	t.Run("Test with empty result", func(t *testing.T) {
		expectedResult := ""
		result := test()
		if result != expectedResult {
			t.Errorf("Expected '%s', but got '%s'", expectedResult, result)
		}
	})
}

func TestMinhaFuncao(t *testing.T) {
	t.Run("Test with id = 7", func(t *testing.T) {
		expectedResult := "OK"
		result := minhaFuncao(7)
		if result != expectedResult {
			t.Errorf("Expected '%s', but got '%s'", expectedResult, result)
		}
	})

	t.Run("Test with id != 7", func(t *testing.T) {
		expectedResult := "OK"
		result := minhaFuncao(5)
		if result != expectedResult {
			t.Errorf("Expected '%s', but got '%s'", expectedResult, result)
		}
	})
}

func TestMainFunction(t *testing.T) {
	t.Run("Test main function", func(t *testing.T) {
		expectedResult := "OK"
		result := captureOutput(func() {
			main()
		})
		if result != expectedResult {
			t.Errorf("Expected '%s', but got '%s'", expectedResult, result)
		}
	})
}

// Função auxiliar para capturar a saída do stdout
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
