package main

import (
	"fmt"
	"strings"
)

// count consecutives english letters into a string and return another string
// input: "aaabccc--@@ddee zzfffz" => output: "a3b1c3d2e2f3"
// input: "aaabccc--@@@ddee ffff ggg xxx zzzz" => output: "a3b1c3d2e2f4g3x3z4"
// input: "a" => "a1"
// input: "abcd" => "a1b1c1d1"
// Time O(n) | Space O(1) where n is the number of characters into the string

func CountOccurences(s string) string {

	if len(s) == 0 {
		return s
	}

	var result strings.Builder
	left, right := 0, 0
	char := s[left]

	for right < len(s) {
		//fmt.Println("char:", string(char), "right:", right, "left:", left, "count:", count)

		if right == len(s)-1 { // stop once we reach the end of the string
			if 97 <= char && char <= 122 { // to consider only english letters
				result.WriteString(fmt.Sprintf("%s%d", string(char), right-left+1))
			}
			break
		}

		if s[right+1] == char {
			right++
			continue
		}

		if 97 <= char && char <= 122 {
			result.WriteString(fmt.Sprintf("%s%d", string(char), right-left+1))
		}

		right++
		left = right
		char = s[left]
	}

	return result.String()
}
