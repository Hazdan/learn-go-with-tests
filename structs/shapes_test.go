package shapes

import "testing"

func TestShape(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	expected := 40.0

	if got != expected {
		t.Errorf("expected %.2f but got %.2f", expected, got)
	}
}
