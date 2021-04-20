package main

import (
	"fmt"
)

func main() {
	tt := []struct {
		array []int
		k int
		output []int
	}{
		{
			array: []int{3, 8, 9, 7, 6},
			k: 1,
			output: []int{6,3,8,9,7},
		},
		{
			array: []int{3, 8, 9,7,6 },
			k: 2,
			output: []int{7,6,3,8,9},
		},
		{
			array: []int{3, 8, 9, 7, 6},
			k: 3,
			output: []int{9,7,6,3,8},
		},
		{
			array: []int{3, 8, 9, 7, 6},
			k: 4,
			output: []int{8,9,7,6,3},
		},
		{
			array: []int{0,0,0},
			k: 1,
			output: []int{0,0,0},
		},
		{
			array: []int{1,2,3,4},
			k: 4,
			output: []int{1,2,3,4},
		},
		{
			array: []int{5, -1000},
			k: 1,
			output: []int{-1000, 5},
		},
		{
			array: []int{5},
			k: 4,
			output: []int{5},
		},


	}

	for _, tc:= range tt {
		res := Solution(tc.array, tc.k)
		if len(res) != len(tc.output) {
			fmt.Printf("inequal: %+v %+v\n", tc.output, res)
		}
		for i := range res {
			if res[i] != tc.output[i] {
				fmt.Printf("expected %+v, got %+v\n", tc.output, res)
				break
			}
 		}
	}
}

func Solution(array []int, k int) []int {
	if len(array) == 0{
		return []int{}
	}
	k = k % len(array)
	res := array
	for n := 0; n < k; n++ {
		res = append([]int{res[len(res)-1]}, res[0:len(res)-1]...)
	}
	return res
}
