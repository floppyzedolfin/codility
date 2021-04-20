package main

import (
	"fmt"
)

func main() {
	tt := []struct {
		a []int
		res int
	} {
		{a: []int{9,3,9,3,9,7,9}, res: 7},
		{a: []int{7}, res: 7},
		{a: []int{9,3,9,3,9,9,9}, res: 9},
	}

	for _, tc := range tt {
		r := Solution(tc.a)
		if tc.res != r {
			fmt.Printf("expected %d, got %d; %+v\n", tc.res, r, tc.a)
		}
	}
}

func Solution(a []int) int {
	m := map[int]struct{}{}
	for _, v := range a {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = struct{}{}
		}
	}
	for k := range m {
		return k
	}
	return -1
}
