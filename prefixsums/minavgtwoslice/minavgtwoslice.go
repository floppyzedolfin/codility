package main

import (
	"fmt"
)

func main() {
	tt := []struct {
		a []int
		r int
	} {
		{[]int{4,2,2,5,1,5,8}, 1},
	}

	for _, tc := range tt {
		r := Solution(tc.a)
		if r != tc.r {
			fmt.Printf("got %d for %+v\n", r, tc)
		}
	}
}

func Solution(a []int) int {
	n := len(a)
	partialSums := make([]int, n)
	s := 0
	p0 := 0
	min := float32(a[0]+a[1])/2.

	for i, v := range a {
		s += v
		partialSums[i] = s
	}

	fmt.Printf("partial sums: %+v\n", partialSums)

	for p := 1; p < n-1; p++ {
		for q := p+1; q < n; q++ {
			m := float32(partialSums[q]-partialSums[p-1])/float32(q-p+1)
			fmt.Printf("p: %d, q: %d, m %f\n", p, q, m)
			if m < min {
				p0 = p
			}
		}
	}

	return p0
}
