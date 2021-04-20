package main

import (
	"fmt"
)

func main() {
	tt := []struct {
		in int
		gap int
	} {
		{in: 9, gap: 2},
		{in: 529, gap: 4},
		{in: 20, gap: 1},
		{in: 15, gap: 0},
		{in: 32, gap: 0},
	}

	for _, tc := range tt {
		fmt.Printf("expected: %d, got: %d, diff: %d\n", tc.gap, Solution(tc.in), tc.gap - Solution(tc.in))
	}
}

func Solution(N int) int {
	max := 0
	count := 0
	for N % 2 == 0 {
		N >>= 1
	}
	for N > 0 {
		switch N%2 {
		case 0:
			count++
		case 1:
			if count > max {
				max = count
			}
			count = 0
		}
		N >>= 1
	}
	return max
}
