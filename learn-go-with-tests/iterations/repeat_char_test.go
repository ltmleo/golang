package iterations

import "testing"

func TestRepeatChar(t *testing.T){
	repeated := repeatChar("a", 5)
	expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected '%s' but got '%s'", expected, repeated)
    }

}

func BenchmarkRepeatChar(b *testing.B) {  //go test -bench=.
    for i := 0; i < b.N; i++ {
        repeatChar("a", 5)
    }
}