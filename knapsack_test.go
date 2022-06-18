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

	if bestVal := GetBestValueRecursion(items, maxWeight); expectBestVal != bestVal {
		t.Errorf("want %d, get %d", expectBestVal, bestVal)
	}
}
