package main

import "testing"

func TestKidsWithCandies(t *testing.T) {
	var testsPositive = []struct {
		in           []int
		candiesExtra int
		want         []bool
	}{
		{
			[]int{2, 3, 5, 1, 3}, 3,
			[]bool{true, true, true, false, true},
		},
		{
			[]int{4, 2, 1, 1, 2}, 1,
			[]bool{true, false, false, false, false},
		},
		{
			[]int{12, 1, 12}, 10,
			[]bool{true, false, true},
		},
	}
	for _, tt := range testsPositive {
		calculated := kidsWithCandies(tt.in, tt.candiesExtra)
		if len(calculated) != len(tt.want) {
			t.Errorf("len expected: %d, len result: %d", len(tt.want), len(calculated))
		}
		for i, want := range tt.want {
			if calculated[i] != want {
				t.Errorf("want: %v, but got: %v", want, calculated[i])
			}
		}
	}
}
