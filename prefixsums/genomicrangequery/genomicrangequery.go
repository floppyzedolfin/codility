package main

import (
	"fmt"
)

func main() {
	tt := []struct {
		s string
		p []int
		q []int
		r []int
	} {
		{"CAGCCTA", []int{2,5,0}, []int{4,5,6}, []int{2,4,1}},
	}

	for _, tc := range tt {
		r := Solution(tc.s, tc.p, tc.q)
		if len(r) != len(tc.r) {
			fmt.Printf("got %+v for %+v\n", r, tc.r)
			break
		}
		for i := range r {
			if r[i] != tc.r[i] {
				fmt.Printf("had %+v for %+v\n", r, tc.r)
			}
		}
	}
}

func Solution(s string, p, q []int) []int {
	r := make([]int, len(p))
outer:
	for i := range p {
		sub := s[p[i]:q[i]+1]
		m := map[int32]struct{}{}
		for _, s := range sub {
			if s == 'A' {
				r[i] = 1
				goto outer
			}
			m[s] = struct{}{}
		}
		switch {
		case m['C'] == struct{}{}:
			r[i] = 2
		case m['G'] == struct{}{}:
			r[i] = 3
		default:
			r[i] = 4
		}
	}

	return r
}