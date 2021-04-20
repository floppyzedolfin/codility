package main

import "fmt"

func main() {
	tt := []struct {
		a []int
		r int
	} {
		{[]int{0,1,0,1,1}, 5},
		{[]int{1,1,0,1,0,1,1}, 5},
		{[]int{0}, 0},
		{[]int{1,0}, 0},
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
	partialSums := make([]int, n+1)
	zeroes := make([]int, 0)
	s := 0

	for i, v := range a {
		switch v {
		case 0:
			zeroes = append(zeroes, i)
		case 1:
			s += v
		}
		partialSums[i+1] = s
	}

	s = 0
	for _, index := range zeroes {
		s += partialSums[n]-partialSums[index]
		if s > 1000000000 {
			return -1
		}
	}
	return s
}
