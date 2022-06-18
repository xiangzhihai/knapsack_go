package knapsack

type Item struct {
	weight int
	value  int
}

// Solve knapsack problem with recursion
func GetBestValueRecursion(items []Item, totalWeight int) int {
	// From Top to Bottom. We think about whether to take each one or not,
	// and what's the best result from either decision according to current
	// left weight.

	// Create mem. This will be a map of a map, where first key is the index
	// of the decision item, while the next key is the total weight left, and
	// result is the best value it can get
	mem := map[int]map[int]int{}

	var getBestVal func(index, totalWeight int) int
	getBestVal = func(index, totalWeight int) int {
		// Base case
		if index == len(items) || totalWeight == 0 {
			return 0
		}

		// Check the best value is already calculated
		if valMap, ok := mem[index]; ok {
			if bestVal, ok := valMap[totalWeight]; ok {
				return bestVal
			}
		} else {
			mem[index] = make(map[int]int)
		}

		// Decide wether to add this item
		var bestVal int
		if totalWeight >= items[index].weight {
			bestVal = getBestVal(index+1, totalWeight-items[index].weight) + items[index].value
		}

		if valueNotPicking := getBestVal(index+1, totalWeight); valueNotPicking > bestVal {
			bestVal = valueNotPicking
		}

		mem[index][totalWeight] = bestVal

		return mem[index][totalWeight]
	}

	return getBestVal(0, totalWeight)
}
