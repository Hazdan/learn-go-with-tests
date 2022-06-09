package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{12.0, 6.0}, 36.0},
		{Circle{10.0}, 62.83185307179586},
	}

	for _, test := range perimeterTests {
		got := test.shape.Perimeter()

		if got != test.expected {
			t.Errorf("expected %g but got %g", test.expected, got)
		}
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, test := range areaTests {
		got := test.shape.Area()

		if got != test.expected {
			t.Errorf("expected %g but got %g", test.expected, got)
		}
	}

}
