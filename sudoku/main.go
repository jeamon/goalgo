package main

// Given a partial complete solution to a sudoku board,
// determine whether the solution is correct.
// Your function should return a boolean.
// {
//     {5, 1, 7, 0, 9, 8, 2, 0, 4},
//     {2, 8, 9, 1, 3, 4, 7, 0, 6},
//     {3, 4, 6, 2, 7, 0, 0, 9, 1},
//     {6, 7, 2, 8, 4, 9, 3, 1, 5},
//     {1, 3, 8, 5, 2, 6, 9, 4, 7},
//     {9, 5, 4, 7, 1, 3, 6, 8, 2},
//     {4, 9, 5, 3, 6, 2, 1, 7, 8},
//     {7, 2, 3, 4, 8, 1, 5, 6, 9},
//     {8, 6, 1, 9, 5, 7, 4, 2, 3},
// }

import (
	"fmt"
)

func IsValidSudoku(sudoku [][]int) bool {

	// validate row
	for _, row := range sudoku {
		rowMap := make(map[int]struct{})
		for _, v := range row {
			if v == 0 {
				continue
			}
			if _, exists := rowMap[v]; exists {
				return false
			} else {
				rowMap[v] = struct{}{}
			}
		}
	}

	// validate colunm
	for j := range 9 {
		colMap := make(map[int]struct{})
		for _, col := range sudoku {
			v := col[j]
			if v == 0 {
				continue
			}
			if _, exists := colMap[v]; exists {
				return false
			} else {
				colMap[v] = struct{}{}
			}
		}
	}

	// validate grid
	for k := 0; k < 3; k++ {
		for k1 := 0; k1 < 3; k1++ {
			m := make(map[int]struct{})
			for i := k * 3; i < k*3+3; i++ {
				for j := k1 * 3; j < k1*3+3; j++ {
					v := sudoku[i][j]
					if v == 0 {
						continue
					}
					if _, exists := m[v]; exists {
						return false
					}

					m[v] = struct{}{}
				}
			}
		}

	}

	return true
}

func main() {

	sudoku := [][]int{
		{5, 1, 7, 0, 9, 8, 2, 0, 4},
		{2, 8, 9, 1, 3, 4, 7, 0, 6},
		{3, 4, 6, 2, 7, 0, 0, 9, 1},
		{6, 7, 2, 8, 4, 9, 3, 1, 5},
		{1, 3, 8, 5, 2, 6, 9, 4, 7},
		{9, 5, 4, 7, 1, 3, 6, 8, 2},
		{4, 9, 5, 3, 6, 2, 1, 7, 8},
		{7, 2, 3, 4, 8, 1, 5, 6, 9},
		{8, 6, 1, 9, 5, 7, 4, 2, 3},
	}

	fmt.Println(IsValidSudoku(sudoku))
}
