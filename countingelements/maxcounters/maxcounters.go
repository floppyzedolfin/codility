package main

import (
	"fmt"
)

func main() {
	tt := []struct {
		n int
		a []int
		out []int
	} {
		{5, []int{3,4,4,6,1,4,4}, []int{3,2,2,4,2}},
		{2, []int{1,3,1,3,2,1}, []int{3,3}},
	}

	for _, tc := range tt {
		r := Solution(tc.n, tc.a)
		if len(r) != len(tc.out) {
			fmt.Printf("inequal lengths %+v %+v\n", r, tc)
			break
		}
		for i := range r {
			if r[i] != tc.out[i] {
				fmt.Printf("inequal values at %d %+v %+v\n", i, r, tc)
				break
			}
		}
	}
}

func Solution(n int, a []int) []int {
	counters := make(map[int]int)
	max := 0
	m :=0

	for i := 0; i < len(a); i++ {
		v := a[i]-1
		if v >= 0 && v < n {
			counters[v]++
			if counters[v] > m {
				m = counters[v]
			}
		} else {
			counters = map[int]int{}
			max += m
			m = 0
		}
		fmt.Printf("max: %d, counters: %+v\n", max, counters)
	}

	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, max + counters[i])
	}

	return res
}