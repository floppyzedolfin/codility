package main

import (
	"fmt"
)

func main() {
	tt := []struct {
		a int
		b int
		k int
		r int
	} {
		{6,11,2,3},
		{5,11,2,3},
		{6,12,2,4},
		{5,12,2,4},
		{11, 345, 17, 20},

	}

	for _ , tc := range tt {
		r := Solution(tc.a, tc.b, tc.k)
		if r != tc.r {
			fmt.Printf("got %d for %+v\n", r, tc)
		}
	}
}

func Solution(a,b,k int) int {
	if b < k {
		if a == 0 {
			return 1
		} else {
			return 0
		}
	}
	s := (a+k-1)/k
	s *= k
	if s > b {
		return 0
	}
	r := (b-s)/k+1

	return r
}