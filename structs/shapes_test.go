package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{12.0, 6.0}, 36.0},
		{Circle{10.0}, 62.83185307179586},
		{Triangle{12, 6}, 31.41640786499874},
	}

	for _, test := range perimeterTests {
		got := test.shape.Perimeter()

		if got != test.expected {
			t.Errorf("%#v expected %g but got %g", test.shape, test.expected, got)
		}
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{"Rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"Circle", Circle{10.0}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()

			if got != test.expected {
				t.Errorf("%#v expected %g but got %g", test.shape, test.expected, got)
			}
		})
	}

}
