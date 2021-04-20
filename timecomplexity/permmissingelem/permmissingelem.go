package main

import (
	"fmt"
)

func main() {
	tt := []struct{
		a []int
		res int
	} {
		{[]int{2,3,1,5}, 4},
		{[]int{1,3}, 2},
	}

	for _, tc := range tt {
		r := Solution(tc.a)
		if r != tc.res {
			fmt.Printf("diff: %d for %+v\n", r, tc)
		}
	}
}

func Solution(a []int) int {
	s := 0
	for _, i := range a {
		s += i
	}
	n := len(a)+1
	return n*(n+1)/2 - s
}