package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("perimeter for rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Perimeter()
		expected := 36.0

		if got != expected {
			t.Errorf("expected %.2f but got %.2f", expected, got)
		}
	})

	t.Run("perimeter for circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		expected := 62.83185307179586

		if got != expected {
			t.Errorf("expected %g but got %g", expected, got)
		}
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()

		got := shape.Area()

		if got != expected {
			t.Errorf("expected %.2f but got %.2f", expected, got)
		}
	}

	t.Run("area for rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		expected := 72.0

		checkArea(t, rectangle, expected)
	})

	t.Run("area for circle", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793

		checkArea(t, circle, expected)
	})
}
