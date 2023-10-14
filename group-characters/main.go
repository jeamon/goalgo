package main

import (
	"strings"
)

// GroupCharacter takes a string s and returns another string which
// contains block(s) of 3 alphanumeric characters.Only the last
// 2 blocks must be formed with 2 alphanumeric characters.
//
// Examples :
// JEROME4AMON2023 returns JER OME 4AM ON2 023
// JEROME 4A--MON202 returns JER OME 4AM ON2 02
// JER OME 4AM ON 02 returns JER OME 4AM ON 02
func GroupCharacters(s string) string {
	var result strings.Builder
	var j int
	for _, r := range s {
		if r == '-' || r == ' ' {
			continue
		}
		result.WriteRune(r)
		j++
		if j%3 == 0 {
			result.WriteRune(' ')
		}
	}

	// remove any automatically added space or newline at the end.
	res := strings.TrimSpace(result.String())
	lg := len(res)

	if lg <= 3 {
		return res
	}

	// last element is at position `(lg -1)`. If we have an empty string before the
	// last element then pick that character before it and rebuids the final string.
	if string(res[lg-2]) == " " {
		res = res[:lg-3] + " " + string(res[lg-3]) + string(res[lg-1])
	}

	return res
}
