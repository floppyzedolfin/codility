package main

import (
	"fmt"
)

func main() {
	tt := []struct{
		a []int
		r int
	} {
		{[]int{2,1,1,2,3,1}, 3},
	}

	for _, tc := range tt {
		r := Solution(tc.a)
		if r != tc.r {
			fmt.Printf("got %d for %+v\n", r, tc)
		}
	}
}

func Solution(a []int) int {
	if len(a) == 0 {
		return 0
	}
	m := map[int]struct{}{}
	for _, v := range a {
		m[v] = struct{}{}
	}
	return len(m)
}