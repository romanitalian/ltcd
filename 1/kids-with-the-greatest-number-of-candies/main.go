package main

import "fmt"

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	res := kidsWithCandies(candies, extraCandies)
	fmt.Println(res)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	mx := max(candies)
	var res []bool
	for _, candy := range candies {
		res = append(res, candy + extraCandies >= mx)
	}
	return res
}

func max(in []int) int {
	l := len(in)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return in[0]
	}
	m := in[0]
	for i := 1; i < l; i++ {
		if in[i] > m {
			m = in[i]
		}
	}
	return m
}