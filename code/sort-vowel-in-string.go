package code

import (
	"fmt"
	"sort"
)

func sortVowels(s string) string {
	vowel := []byte{}
	isVowel := [256]bool{}

	for _, ch := range []byte("aeiouAEIOU") {
		isVowel[ch] = true
	}

	for _, ch := range s {
		if isVowel[ch] {
			vowel = append(vowel, byte(ch))
		}
	}

	sort.Slice(vowel, func(i, j int) bool { return vowel[i] < vowel[j] })

	res := []byte(s)
	j := 0

	for i := range res {
		if isVowel[res[i]] {
			res[i] = vowel[j]
			j++
		}
	}
	return string(res)
}

func Test13() {
	tests := []string{"LEetcOde", "lYmpH"}

	for _, val := range tests {
		fmt.Println(sortVowels(val))
	}
}
