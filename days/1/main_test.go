package main

import (
  "testing"
)

func TestSecondPart(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
		firstSlice := []int{3, 4, 2, 1, 3, 3}
		secondSlice := []int{4, 3, 5, 3, 9, 3}

		got := SolutionSecondPart(firstSlice, secondSlice)
		want := 31

		if got != want {
			t.Errorf("want %d got %d", want, got)
		}
	})
}
