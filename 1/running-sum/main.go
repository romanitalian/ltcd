package main

import "fmt"

func main() {
	nums := []int{-1,-2,-30,4}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {
	for i, _ := range nums {
		if i == 0 {
			continue
		}
		nums[i] += nums[i-1]
	}
	return nums
}