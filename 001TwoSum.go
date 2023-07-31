package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// Time: O(n^2)
	// Space: O(1)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	var nums []int
	var target int

	nums, target = []int{17, 2, 3, 8, 19, 7, 11, 15}, 9
	fmt.Println(nums, target, "\nSolution:", twoSum(nums, target))
	fmt.Println()

	nums, target = []int{-3, 4, 3, 90}, 0
	fmt.Println(nums, target, "\nSolution:", twoSum(nums, target))
	fmt.Println()
}
