package code

import (
	"fmt"
	"math"
	"sort"
)

type q18 struct {
	nums []int
	k    int
}

func maxFrequency(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	left := 0
	ans := 0
	cur := 0

	for right := range nums {
		target := nums[right]
		cur += target

		for (right-left+1)*target-cur > k {
			cur -= nums[left]
			left += 1
		}
		ans = int(math.Max(float64(ans), float64(right-left+1)))

	}
	return ans

}

func Test18() {
	tests := []q18{
		{[]int{1, 2, 4}, 5},
		{[]int{1, 4, 8, 13}, 5},
		{[]int{3, 9, 6}, 3},
	}

	for _, test := range tests {
		fmt.Println(maxFrequency(test.nums, test.k))
	}
}
