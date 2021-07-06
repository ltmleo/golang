package hello
import (
	"testing" 
	// . "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

func TestHelloWorld(t *testing.T){
	t.Run("Hello World Leo", func(t *testing.T){
		got := Hello("Leo")
		expected := "Hello World Leo"
		if got != expected  {
			t.Errorf("expected: %s, got: %s", expected, got)
		}
	})
	t.Run("Hello World Lais", func(t *testing.T){
		got := Hello("Lais")
		expected := "Hello World Lais"
		if got != expected  {
			t.Errorf("expected: %s, got: %s", expected, got)
		}
	})
}

// var _ = Describe("Hello", func() {
// 		Describe("Saying Hello", func() {
// 			It("to Leo", func() {
// 				got := Hello("Leo")
// 				Expect(got).To(Equal("Hello World Leo"))
// 			})
// 		})
//     })