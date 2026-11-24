package main

import "fmt"

// MinCostClimbing returns the cheapest total cost to pass the staircase:
// pay cost[i] to stand on step i, hop 1 or 2, start at step 0 or 1, and
// finish beyond the last step.
func MinCostClimbing(cost []int) int {
	if len(cost) < 2 {
		return 0 // one step or none can be skipped over entirely
	}
	a, b := cost[0], cost[1] // cheapest total ending ON steps i-2, i-1
	for i := 2; i < len(cost); i++ {
		m := a
		if b < m {
			m = b
		}
		a, b = b, cost[i]+m
	}
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(MinCostClimbing([]int{10, 15, 20}))
	fmt.Println(MinCostClimbing([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
