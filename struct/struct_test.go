package main

import "testing"

// Testing Perimeter
func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		got := Perimeter(rectangle)
		want := 60.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

// Testing Area
func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := Area(rectangle)
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %2.f", got, want)
		}
	})
}
