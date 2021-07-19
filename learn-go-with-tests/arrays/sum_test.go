package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){
	t.Run("collection of 4 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		sum := Sum(numbers)
		expected := 10
		if sum != expected {
			t.Errorf("Sum %d is different than expected %d", sum, expected)
		}
	})
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Sum(numbers)
		expected := 15
		if sum != expected {
			t.Errorf("Sum %d is different than expected %d", sum, expected)
		}
	})
}

func TestSumAll(t *testing.T){
	t.Run("One collection of 4 numbers", func(t *testing.T) {
		sum := SumAll([]int{1, 2, 3, 4})
		expected := []int{10}
		if !reflect.DeepEqual(sum, expected) {
			t.Errorf("Sum %v is different than expected %v", sum, expected)
		}
	})
	t.Run("Two collection of numbers", func(t *testing.T) {
		sum := SumAll([]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5})
		expected := []int{10, 15}
		if !reflect.DeepEqual(sum, expected) {
			t.Errorf("Sum %v is different than expected %v", sum, expected)
		}
	})
}
// go test -cover