package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoveToLeft(t *testing.T) {
	cases := []struct {
		input  []int
		places int
		want   []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			places: 3,
			want:   []int{4, 5, 6, 7, 8, 1, 2, 3},
		},
		{
			input:  []int{4, 5, 6, 7, 8, 1, 2, 3},
			places: 3,
			want:   []int{7, 8, 1, 2, 3, 4, 5, 6},
		},
		{
			input:  []int{1},
			places: 3,
			want:   []int{1},
		},
		{
			input:  []int{1, 2},
			places: 3,
			want:   []int{2, 1},
		},
		{
			input:  []int{},
			places: 4,
			want:   []int{},
		},
		{
			input:  nil,
			places: 4,
			want:   nil,
		},
	}

	for n, tc := range cases {
		t.Run(fmt.Sprint("case:", n+1), func(t *testing.T) {
			got := MoveToLeft(tc.input, tc.places)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v but want: %v", got, tc.want)
			}
		})
	}
}

func TestMoveToRight(t *testing.T) {
	cases := []struct {
		input  []int
		places int
		want   []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			places: 3,
			want:   []int{6, 7, 8, 1, 2, 3, 4, 5},
		},
		{
			input:  []int{6, 7, 8, 1, 2, 3, 2, 5},
			places: 3,
			want:   []int{3, 2, 5, 6, 7, 8, 1, 2},
		},
		{
			input:  []int{1},
			places: 3,
			want:   []int{1},
		},
		{
			input:  []int{1, 2},
			places: 3,
			want:   []int{2, 1},
		},
		{
			input:  []int{},
			places: 4,
			want:   []int{},
		},
		{
			input:  nil,
			places: 4,
			want:   nil,
		},
	}

	for n, tc := range cases {
		t.Run(fmt.Sprint("case:", n+1), func(t *testing.T) {
			got := MoveToRight(tc.input, tc.places)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v but want: %v", got, tc.want)
			}
		})
	}
}
