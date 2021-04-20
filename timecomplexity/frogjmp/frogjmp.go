package main

import (
	"fmt"
)

func main() {
	tt := []struct {
		x int
		y int
		d int
		res int
	} {
		{10, 85, 30, 3},
		{10, 100, 30, 3},
		{10, 101, 30, 4},
		{10, 10, 30, 0},
		{10, 100, 30000, 1},
	}

	for _, tc := range tt {
		r := Solution(tc.x, tc.y, tc.d)
		if r != tc.res {
			fmt.Printf("expected %d, got %d, %+v\n", r, tc.res, tc)
		}
	}
}


func Solution(x,y,d int) int {
	distance := y -x
	return (distance + d - 1) /d

}