package main

import (
	"fmt"
	"testing"
)

func TestGroupCharacters(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"", ""},
		{" ", ""},
		{"J", "J"},
		{"JE", "JE"},
		{"JER", "JER"},
		{"JERO", "JE RO"},
		{"JEROME4AMON2023", "JER OME 4AM ON2 023"},
		{"JEROME 4A--MON202", "JER OME 4AM ON2 02"},
		{"JER OME 4AM ON 02", "JER OME 4AM ON 02"},
		{"JE---RO ME 4A--MON2   0 -- 2", "JER OME 4AM ON2 02"},
	}

	for n, tc := range cases {
		t.Run(fmt.Sprint("case:", n+1), func(t *testing.T) {
			got := GroupCharacters(tc.input)
			if got != tc.want {
				t.Errorf("got: %q but want: %q", got, tc.want)
			}
		})
	}
}

/*
$ go test -v group-characters/*
=== RUN   TestGroupCharacters
=== RUN   TestGroupCharacters/case:1
=== RUN   TestGroupCharacters/case:2
=== RUN   TestGroupCharacters/case:3
=== RUN   TestGroupCharacters/case:4
=== RUN   TestGroupCharacters/case:5
=== RUN   TestGroupCharacters/case:6
=== RUN   TestGroupCharacters/case:7
=== RUN   TestGroupCharacters/case:8
=== RUN   TestGroupCharacters/case:9
=== RUN   TestGroupCharacters/case:10
--- PASS: TestGroupCharacters (0.00s)
    --- PASS: TestGroupCharacters/case:1 (0.00s)
    --- PASS: TestGroupCharacters/case:2 (0.00s)
    --- PASS: TestGroupCharacters/case:3 (0.00s)
    --- PASS: TestGroupCharacters/case:4 (0.00s)
    --- PASS: TestGroupCharacters/case:5 (0.00s)
    --- PASS: TestGroupCharacters/case:6 (0.00s)
    --- PASS: TestGroupCharacters/case:7 (0.00s)
    --- PASS: TestGroupCharacters/case:8 (0.00s)
    --- PASS: TestGroupCharacters/case:9 (0.00s)
    --- PASS: TestGroupCharacters/case:10 (0.00s)
PASS
ok      command-line-arguments  0.136s
*/
