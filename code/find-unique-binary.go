package code

import "fmt"

func rec(s string, length int, store *[]string) {
	if len(s) == length {
		*store = append(*store, s)
	}
	if len(s) > length {
		return
	}

	rec(s+"0", length, store)
	rec(s+"1", length, store)
}

func findDifferentBinayString(nums []string) string {
	store := []string{}
	length := len(nums[0])

	rec("", length, &store)
	isPresent := func(str string) bool {
		for _, val := range nums {
			if val == str {
				return false
			}
		}
		return true
	}

	for _, val := range store {
		if isPresent(val) {
			return val
		}
	}
	return ""
}

func Test16() {
	tests := [][]string{
		{"10", "10"},
		{"00", "01"},
		{"111", "011", "001"},
	}

	for _, val := range tests {
		fmt.Println(findDifferentBinayString(val))
	}
}
