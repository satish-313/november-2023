package code

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	maxSum := -1
	sum := 0
	left := 0
	right := len(nums) - 1

	for left < right {
		sum = nums[left] + nums[right]
		if maxSum < sum {
			maxSum = sum
		}
		left += 1
		right -= 1
	}

	return maxSum
}

func Test17() {
	tests := [][]int{
		{3, 5, 2, 3},
		{3, 5, 4, 2, 4, 6},
	}

	for _, val := range tests {
		fmt.Println(minPairSum(val))
	}

}
