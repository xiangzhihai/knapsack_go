package knapsack

import "testing"

func TestGetBestValueRecursion(t *testing.T) {
	items := []Item{
		{1, 1},
		{2, 1},
		{4, 5},
		{2, 3},
		{1, 2},
		{5, 2},
	}
	maxWeight := 5
	expectBestVal := 7

	// Check recursion
	if bestVal := GetBestValueRecursion(items, maxWeight); expectBestVal != bestVal {
		t.Errorf("recursion wants %d, get %d", expectBestVal, bestVal)
	}

	// Check DP
	if bestVal := GetBestValueDP(items, maxWeight); expectBestVal != bestVal {
		t.Errorf("dynamic programming wants %d, get %d", expectBestVal, bestVal)
	}
}
