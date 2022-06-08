package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	}

	t.Run("sum all tails with an empty slice", func(t *testing.T) {
		array1 := []int{1, 2, 3, 4, 5}
		array2 := []int{}

		got := SumAllTails(array1, array2)
		expected := []int{14, 0}

		checkSums(t, got, expected)
	})
	t.Run("sum all tails", func(t *testing.T) {
		array1 := []int{1, 2, 3, 4, 5}
		array2 := []int{2, 3}

		got := SumAllTails(array1, array2)
		expected := []int{14, 3}

		checkSums(t, got, expected)

	})
}

func TestSumAll(t *testing.T) {
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{2, 3}

	got := SumAll(array1, array2)
	expected := []int{15, 5}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v but got %v, given %v and %v", expected, got, array1, array2)
	}
}
func TestSum(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, expected int, numbers []int) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %d but got %d, given %v", expected, got, numbers)
		}
	}

	t.Run("sum array with any numbers", func(t *testing.T) {
		numbers := []int{1, 3, 4}

		got := Sum(numbers)
		expected := 8

		assertCorrectMessage(t, got, expected, numbers)

	})
}
