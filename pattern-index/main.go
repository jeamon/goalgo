package main

// indexOf(pattern, text): returns the index of the first occurrence of pattern in text, or -1 if none occurs.
// indexOf("ab", "babc") = 1
// indexOf("ab", "bacb") = -1
// Time: O(n*m) Space: O(1)

func IndexOf(pattern, text string) int {
	m, n := len(pattern), len(text)

	switch {
	case m == 0:
		return 0
	case m == 1:
		if n > 0 && pattern == string(text[0]) {
			return 0
		}
		return -1
	case m == n:
		if pattern == text {
			return 0
		}
		return -1
	case m > n:
		return -1

	}

	for i := range n {
		if text[i] == pattern[0] {
			end := true
			for j := 1; j < m; j++ {
				if text[i+j] != pattern[j] {
					end = false
				}
			}
			if end {
				return i
			}
		}
	}
	return -1
}
