package structs

import "testing"

var rectangle = Rectangle{10.0, 10.0}
var circle = Circle{10}

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {

    areaTests := []struct {
		name string
        shape Shape
        hasArea  float64
    }{
        {"Rectangle", Rectangle{12, 6}, 72.0}, // not clearly
        {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.15000000000003},
        {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
    }

    for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
    }

}

// Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer
// Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)
// Adding methods so you can add functionality to your data types and so you can implement interfaces
// Table driven tests to make your assertions clearer and your test suites easier to extend & maintain