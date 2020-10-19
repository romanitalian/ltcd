package main

import (
	"testing"
)

func TestRunningSum(t *testing.T) {
	var testsPositive = []struct {
		in   []int
		want []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 3, 6, 10},
		},
		{
			[]int{1, 1, 1, 1, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{3, 1, 2, 10, 1},
			[]int{3, 4, 6, 16, 17},
		},
		{
			[]int{-1, 2, 3, 4},
			[]int{-1, 1, 4, 8},
		},
		{
			[]int{-1, -2, -30, 4},
			[]int{-1, -3, -33, -29},
		},
	}

	for _, tt := range testsPositive {
		runningSumValues := runningSum(tt.in)
		if len(runningSumValues) != len(tt.want) {
			t.Errorf("len expected: %d, len result: %d", len(tt.want), len(runningSumValues))
		}
		for i, want := range tt.want {
			if runningSumValues[i] != want {
				t.Errorf("want: %d, but got: %d", want, runningSumValues[i])
			}
		}
	}
}
