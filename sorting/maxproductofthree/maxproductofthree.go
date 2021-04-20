package main

import (
	"fmt"
	"sort"
)

func main() {
	tt := []struct {
		a []int
		r int
	}{
		{[]int{-3, 1, 2, -2, 5, 6}, 60},
		{[]int{-3, -2, -1, 0}, 0},
		{[]int{-3, -2, -1}, -6},
		{[]int{-1, -2, -3, -4}, -6},
		{[]int{-3, -2, -1, 0, 7}, 42},
		{[]int{-3, -2, -1, 0, 7, 8}, 48},
		{[]int{-3, -2, -1, 0, 4, 5, 6}, 120},
	}

	for _, tc := range tt {
		r := Solution(tc.a)
		if r != tc.r {
			fmt.Printf("got %d for %+v\n", r, tc)
		}
	}
}

func Solution(a []int) int {
	sort.Ints(a)
	negs := a[0:2]
	pos := a[len(a)-3:]

	n := negs[0] * negs[1]
	p := pos[0] * pos[1]
	if n < p || pos[2] < 0 {
		return p * pos[2]
	} else {
		return n * pos[2]
	}
}
