package main

import (
	"fmt"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{
			input: "",
			want:  "",
		},
		{
			input: "a",
			want:  "a1",
		},
		{
			input: "ab cd e",
			want:  "a1b1c1d1e1",
		},
		{
			input: "aabcddd @@e -- zzz",
			want:  "a2b1c1d3e1z3",
		},
		{
			input: "#  xx | zz | aa | v4",
			want:  "x2z2a2v1",
		},
	}

	for n, tc := range cases {
		t.Run(fmt.Sprint("case:", n+1), func(t *testing.T) {
			got := CountOccurences(tc.input)
			if got != tc.want {
				t.Errorf("got: %v but want: %v", got, tc.want)
			}
		})
	}
}

/*

go test -timeout 30s -run ^TestSuggestedProducts$ github.com/jeamon/goalgo/count-occurences -count=1 -v
=== RUN   TestSuggestedProducts
=== RUN   TestSuggestedProducts/case:1
=== RUN   TestSuggestedProducts/case:2
=== RUN   TestSuggestedProducts/case:3
=== RUN   TestSuggestedProducts/case:4
=== RUN   TestSuggestedProducts/case:5
--- PASS: TestSuggestedProducts (0.00s)
    --- PASS: TestSuggestedProducts/case:1 (0.00s)
    --- PASS: TestSuggestedProducts/case:2 (0.00s)
    --- PASS: TestSuggestedProducts/case:3 (0.00s)
    --- PASS: TestSuggestedProducts/case:4 (0.00s)
    --- PASS: TestSuggestedProducts/case:5 (0.00s)
PASS
ok      github.com/jeamon/goalgo/count-occurences       0.167s

*/
