package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 3)
	expected := "aaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
