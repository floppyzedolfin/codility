package main

import (
	"fmt"
)

func main() {
	tt := []struct{
		x int
		a []int
		r int
	} {
		{5, []int{1,3,1,4,2,3,5,4}, 6},
		{5, []int{1,3,1,4,2,3,4,4}, -1},

	}

	for _, tc := range tt {
		r := Solution(tc.x, tc.a)
		if r != tc.r {
			fmt.Printf("got %d for %+v\n", r, tc)
		}
	}
}

func Solution(x int, a []int) int {
	leaves := make(map[int]struct{})

	for i, v := range a {
		leaves[v] = struct{}{}
		if len(leaves) == x {
			return i
		}
	}
	return -1

}