package main

import (
	"fmt"
	"testing"
)

func TestIndexOf(t *testing.T) {
	cases := []struct {
		text    string
		pattern string
		want    int
	}{
		{
			text:    "",
			pattern: "",
			want:    0,
		},
		{
			text:    "",
			pattern: "a",
			want:    -1,
		},
		{
			text:    "a",
			pattern: "",
			want:    0,
		},
		{
			text:    "a",
			pattern: "a",
			want:    0,
		},
		{
			text:    "aba",
			pattern: "a",
			want:    0,
		},
		{
			text:    "babc",
			pattern: "ab",
			want:    1,
		},
		{
			text:    "bacb",
			pattern: "ab",
			want:    -1,
		},
		{
			text:    "amon abou jerome",
			pattern: "rome",
			want:    12,
		},
	}

	for n, tc := range cases {
		t.Run(fmt.Sprint("case:", n+1), func(t *testing.T) {
			got := IndexOf(tc.pattern, tc.text)
			if got != tc.want {
				t.Errorf("got: %v but want: %v", got, tc.want)
			}
		})
	}
}

/*

go test -timeout 30s -run ^TestIndexOf$ github.com/jeamon/goalgo/pattern-index -count=1 -v
=== RUN   TestIndexOf
=== RUN   TestIndexOf/case:1
=== RUN   TestIndexOf/case:2
=== RUN   TestIndexOf/case:3
=== RUN   TestIndexOf/case:4
=== RUN   TestIndexOf/case:5
=== RUN   TestIndexOf/case:6
=== RUN   TestIndexOf/case:7
=== RUN   TestIndexOf/case:8
--- PASS: TestIndexOf (0.00s)
    --- PASS: TestIndexOf/case:1 (0.00s)
    --- PASS: TestIndexOf/case:2 (0.00s)
    --- PASS: TestIndexOf/case:3 (0.00s)
    --- PASS: TestIndexOf/case:4 (0.00s)
    --- PASS: TestIndexOf/case:5 (0.00s)
    --- PASS: TestIndexOf/case:6 (0.00s)
    --- PASS: TestIndexOf/case:7 (0.00s)
    --- PASS: TestIndexOf/case:8 (0.00s)
PASS
ok      github.com/jeamon/goalgo/pattern-index  0.267s

*/
