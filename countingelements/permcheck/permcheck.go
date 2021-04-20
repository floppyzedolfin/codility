package main

import (
	"fmt"
)

func main() {
	tt := []struct {
		a []int
		r int
	} {
		{[]int{4,1,3,2}, 1},
		{[]int{4,1,3}, 0},
		{[]int{-1, 0, 1}, 0},


	}

	for _, tc := range tt {
		r := Solution(tc.a)
		if r != tc.r {
			fmt.Printf("got %d for %+v\n", r, tc)
		}
	}

}

func Solution(a []int) int {
	var s int64
	m := make(map[int]struct{})

	for _, v := range a {
		if _, ok := m[v]; ok {
			return 0
		}

		m[v] = struct{}{}
		s += int64(v)
	}
	n := int64(len(a))
	if int64(len(m)) != n {
		return 0
	}

	if s == n*(n+1) / 2 {
		return 1
	}
	return 0
}