package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	cases := []struct {
		input  []string
		search string
		max    int
		want   [][]string
	}{
		{
			input:  nil,
			search: "word",
			max:    1,
			want: [][]string{
				{},
				{},
				{},
				{},
			},
		},
		{
			input:  []string{},
			search: "word",
			max:    1,
			want: [][]string{
				{},
				{},
				{},
				{},
			},
		},
		{
			input:  []string{"word", "wording", "words", "world"},
			search: "word",
			max:    0,
			want: [][]string{
				{},
				{},
				{},
				{},
			},
		},
		{
			input:  []string{},
			search: "",
			max:    1,
			want:   nil,
		},
		{
			input:  []string{"word", "wording", "words", "world"},
			search: "word",
			max:    3,
			want: [][]string{
				{"word", "wording", "words"},
				{"word", "wording", "words"},
				{"word", "wording", "words"},
				{"word", "wording", "words"},
			},
		},
		{
			input:  []string{"word", "wording", "words", "world"},
			search: "words",
			max:    3,
			want: [][]string{
				{"word", "wording", "words"},
				{"word", "wording", "words"},
				{"word", "wording", "words"},
				{"word", "wording", "words"},
				{"words"},
			},
		},
		{
			input:  []string{"word", "woodies", "woodin", "wording", "woodbox", "words", "wood", "world", "woods"},
			search: "woods",
			max:    7,
			want: [][]string{
				{"wood", "woodbox", "woodies", "woodin", "woods", "word", "wording"},
				{"wood", "woodbox", "woodies", "woodin", "woods", "word", "wording"},
				{"wood", "woodbox", "woodies", "woodin", "woods"},
				{"wood", "woodbox", "woodies", "woodin", "woods"},
				{"woods"},
			},
		},
	}

	for n, tc := range cases {
		t.Run(fmt.Sprint("case:", n+1), func(t *testing.T) {
			got := SuggestedProducts(tc.input, tc.search, tc.max)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v but want: %v", got, tc.want)
			}
		})
	}
}

/*

go test -timeout 30s -run ^TestSuggestedProducts$ github.com/jeamon/goalgo/search-suggestions -count=1 -v
=== RUN   TestSuggestedProducts
=== RUN   TestSuggestedProducts/case:1
=== RUN   TestSuggestedProducts/case:2
=== RUN   TestSuggestedProducts/case:3
=== RUN   TestSuggestedProducts/case:4
=== RUN   TestSuggestedProducts/case:5
=== RUN   TestSuggestedProducts/case:6
=== RUN   TestSuggestedProducts/case:7
--- PASS: TestSuggestedProducts (0.00s)
    --- PASS: TestSuggestedProducts/case:1 (0.00s)
    --- PASS: TestSuggestedProducts/case:2 (0.00s)
    --- PASS: TestSuggestedProducts/case:3 (0.00s)
    --- PASS: TestSuggestedProducts/case:4 (0.00s)
    --- PASS: TestSuggestedProducts/case:5 (0.00s)
    --- PASS: TestSuggestedProducts/case:6 (0.00s)
    --- PASS: TestSuggestedProducts/case:7 (0.00s)
PASS
ok      github.com/jeamon/goalgo/search-suggestions     0.275s

*/

func TestTrieSuggestedProducts(t *testing.T) {
	cases := []struct {
		input  []string
		search string
		max    int
		want   [][]string
	}{
		{
			input:  nil,
			search: "word",
			max:    1,
			want: [][]string{
				{},
				{},
				{},
				{},
			},
		},
		{
			input:  []string{},
			search: "word",
			max:    1,
			want: [][]string{
				{},
				{},
				{},
				{},
			},
		},
		{
			input:  []string{"word", "wording", "words", "world"},
			search: "word",
			max:    0,
			want: [][]string{
				{},
				{},
				{},
				{},
			},
		},
		{
			input:  []string{},
			search: "",
			max:    1,
			want:   [][]string{},
		},
		{
			input:  []string{"word", "wording", "words", "world"},
			search: "word",
			max:    3,
			want: [][]string{
				{"word", "wording", "words"},
				{"word", "wording", "words"},
				{"word", "wording", "words"},
				{"word", "wording", "words"},
			},
		},
		{
			input:  []string{"word", "wording", "words", "world"},
			search: "words",
			max:    3,
			want: [][]string{
				{"word", "wording", "words"},
				{"word", "wording", "words"},
				{"word", "wording", "words"},
				{"word", "wording", "words"},
				{"words"},
			},
		},
		{
			input:  []string{"word", "woodies", "woodin", "wording", "woodbox", "words", "wood", "world", "woods"},
			search: "woods",
			max:    7,
			want: [][]string{
				{"wood", "woodbox", "woodies", "woodin", "woods", "word", "wording"},
				{"wood", "woodbox", "woodies", "woodin", "woods", "word", "wording"},
				{"wood", "woodbox", "woodies", "woodin", "woods"},
				{"wood", "woodbox", "woodies", "woodin", "woods"},
				{"woods"},
			},
		},
	}

	for n, tc := range cases {
		t.Run(fmt.Sprint("case:", n+1), func(t *testing.T) {
			got := TrieSuggestedProducts(tc.input, tc.search, tc.max)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v but want: %v", got, tc.want)
			}
		})
	}
}

/*

go test -timeout 30s -run ^TestTrieSuggestedProducts$ github.com/jeamon/goalgo/search-suggestions -count=1 -v
=== RUN   TestTrieSuggestedProducts
=== RUN   TestTrieSuggestedProducts/case:1
=== RUN   TestTrieSuggestedProducts/case:2
=== RUN   TestTrieSuggestedProducts/case:3
=== RUN   TestTrieSuggestedProducts/case:4
=== RUN   TestTrieSuggestedProducts/case:5
=== RUN   TestTrieSuggestedProducts/case:6
=== RUN   TestTrieSuggestedProducts/case:7
--- PASS: TestTrieSuggestedProducts (0.00s)
    --- PASS: TestTrieSuggestedProducts/case:1 (0.00s)
    --- PASS: TestTrieSuggestedProducts/case:2 (0.00s)
    --- PASS: TestTrieSuggestedProducts/case:3 (0.00s)
    --- PASS: TestTrieSuggestedProducts/case:4 (0.00s)
    --- PASS: TestTrieSuggestedProducts/case:5 (0.00s)
    --- PASS: TestTrieSuggestedProducts/case:6 (0.00s)
    --- PASS: TestTrieSuggestedProducts/case:7 (0.00s)
PASS
ok      github.com/jeamon/goalgo/search-suggestions     0.278s

*/
