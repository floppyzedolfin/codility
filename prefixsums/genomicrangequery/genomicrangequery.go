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
	}{
		{"CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}, []int{2, 4, 1}},
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

	a, c, g, t := 0, 0, 0, 0
	as := make([]int, len(s)+1)
	cs := make([]int, len(s)+1)
	gs := make([]int, len(s)+1)
	ts := make([]int, len(s)+1)

	for i, char := range s {
		switch char {
		case 'A':
			a++
		case 'C':
			c++
		case 'G':
			g++
		case 'T':
			t++
		}
		as[i+1] = a
		cs[i+1] = c
		gs[i+1] = g
		ts[i+1] = t
	}

	for i := range p {
		switch {
		case as[q[i]+1]-as[p[i]] > 0:
			r[i] = 1
		case cs[q[i]+1]-cs[p[i]] > 0:
			r[i] = 2
		case gs[q[i]+1]-gs[p[i]] > 0:
			r[i] = 3
		case ts[q[i]+1]-ts[p[i]] > 0:
			r[i] = 4
		default:
			r[i] = -1
		}
	}

	return r
}
