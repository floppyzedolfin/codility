package main

import (
	"fmt"
)

func main() {
	tt := []struct {
		a []int
		res int
	} {
		{[]int{3,1,2,4,3}, 1},
		{[]int{3,1}, 2},
		{[]int{3,3,6}, 0},
		{[]int{-3,-1,-2,-4,-3}, 1},

	}
	for _, tc := range tt {
		r := Solution(tc.a)
		if r != tc.res {
			fmt.Printf("got %v for %+v\n", r, tc.res)
		}
	}
}

func Solution(a []int) int {
	partialSums := make([]int, 0)
	max := abs(a[0])
	s := 0
	n := len(a)
	for _, v := range a {
		if abs(v) >  max {
			max = abs(v)
		}
		s += v
		partialSums = append(partialSums, s)
	}

	min := n * max

	for i := 1; i < n; i++ {
		lhs := partialSums[i-1]
		rhs := partialSums[n-1] - partialSums[i-1]
		if d := abs(lhs - rhs); d < min {
			min = d
			if min == 0 {
				return min
			}
		}
	}

	return min
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}