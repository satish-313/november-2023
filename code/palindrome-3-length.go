package code

import "fmt"

func countPalindromicSubsequence(s string) int {
	m := [26]byte{}
	arr := [26]bool{}

	for _, val := range s {
		m[val-97] = byte(val)
	}

	ans := 0

	for _, ch := range m {
		if ch == 0 {
			continue
		}
		l := -1
		r := 0

		for idx, val := range s {
			if byte(val) == ch {
				if l == -1 {
					l = idx
				}
				r = idx
			}
		}

		for i := l + 1; i < r; i++ {
			arr[s[i]-97] = true
		}
		for i := 0; i < 26; i++ {
			if arr[i] {
				ans += 1
				arr[i] = false
			}
		}
	}

	return ans
}

func Test14() {
	tests := []string{"aabca", "adcz", "bbcbaba"}
	for _, test := range tests {
		fmt.Println(countPalindromicSubsequence(test))
	}
}
