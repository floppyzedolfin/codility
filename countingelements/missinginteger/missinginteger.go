package main

import (
	"fmt"
	"sort"
)

func main() {
	tt := []struct {
		a []int
		r int
	} {
		{[]int{1,3,6,4,1,2}, 5},
		{[]int{1,2,3}, 4},
		{[]int{-1,-3}, 1},
		{[]int{-1,1}, 2},
	}

	for _, tc := range tt {
		r := Solution(tc.a)
		if r != tc.r {
			fmt.Printf("got %d for %+v\n", r, tc)
		}
	}
}

func Solution(a []int) int {
	seen := map[int]struct{}{}
	for _, v := range a {
		if v > 0 {
			seen[v] = struct{}{}
		}
	}
	seenK := make([]int, 0)
	for k := range seen {
		seenK = append(seenK, k)
	}
	sort.Ints(seenK)
	for i, k := range seenK {
		if i +1 != k {
			return i+1
		}
	}
	return len(seenK)+1
}