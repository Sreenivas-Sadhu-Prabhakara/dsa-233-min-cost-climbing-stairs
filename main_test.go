package main

import "testing"

func TestMinCostClimbing(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{[]int{5}, 0},
		{[]int{}, 0},
		{[]int{1, 2}, 1},
		{[]int{0, 0, 0, 1}, 0},
	}
	for _, c := range cases {
		if got := MinCostClimbing(c.in); got != c.want {
			t.Errorf("MinCostClimbing(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
