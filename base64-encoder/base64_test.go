package base64

import (
	"testing"
)

var testCases = []struct {
	testName string
	stringToBeEncoded string
	stringEncoded string

}{
	{"String", "Leonardo", "TGVvbmFyZG8="},
	{"Integer", "123456", "MTIzNDU2"},
	{"Special Characters", "!@#$%", "IUAjJCU="},
	{"Multiple Characters", "Hello World!! 123 #$ :)", "SGVsbG8gV29ybGQhISAxMjMgIyQgOik="},
	{"Empty String", "", ""},
}
var testErrorCases = []struct {
	testName string
	stringToBeDecoded string
	decodedError error
}{
	{"Not Base64 format", "this_is_not_base64", invalidStringError},
}

func TestEncode(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			assert(t, tt.stringToBeEncoded, tt.stringEncoded, Encode(tt.stringToBeEncoded))
		})
	}
}
func TestDecode(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			decodedString, _ := Decode(tt.stringEncoded)
			assert(t, tt.stringEncoded, tt.stringToBeEncoded, decodedString)

		})
	}
}
func TestDecodeErrors(t *testing.T) {
	for _, tt := range testErrorCases {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := Decode(tt.stringToBeDecoded)
			assertError(t, err, tt.decodedError)
		})
	}
}

func assert(t testing.TB, stringToBeEncoded string, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Encode(%q) = %q, want %q", stringToBeEncoded, got, want)
	}
}

func assertError (t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}