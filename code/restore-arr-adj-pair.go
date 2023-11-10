package code

import (
	"fmt"
	"math"
)

func restoreArray(adjacentPairs [][]int) []int {
	res := []int{}

	m := make(map[int][]int)
	prev := math.MaxInt32
	curr := 0

	for _, val := range adjacentPairs {
		a, b := val[0], val[1]
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}

	for i, val := range m {
		if len(val) == 1 {
			res = append(res, i)
			curr = i
			break
		}
	}

	for len(res) < len(adjacentPairs)+1 {
		for _, val := range m[curr] {
			if val != prev {
				prev = curr
				curr = val
				res = append(res, val)
				break
			}
		}
	}

	return res
}

type test10 struct {
	q [][]int
	a []int
}

func Test10() {
	tests := []test10{
		{[][]int{{2, 1}, {3, 4}, {3, 2}}, []int{1, 2, 3, 4}},
		{[][]int{{4, -2}, {1, 4}, {-3, 1}}, []int{-2, -4, 1, -1}},
		{[][]int{{10000, -10000}}, []int{10000, -10000}},
	}

	for _, val := range tests {
		fmt.Printf("Q, %v and A %v my ans := %v \n", val.q, val.a, restoreArray(val.q))
	}
}
