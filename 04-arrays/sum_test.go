package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Expected %d, got %d, given %v", want, got, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("Expected %d, got %d, given %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	numbers := [][]int{[]int{1, 2, 3}, []int{2, 4, 6}}
	got := SumAll(numbers...)
	want := []int{6, 12}

	if reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v, with %v", want, got, numbers)
	}
}
