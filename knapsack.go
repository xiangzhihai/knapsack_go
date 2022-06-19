package knapsack

type Item struct {
	weight int
	value  int
}

// Solve knapsack problem with recursion
func GetBestValueRecursion(items []Item, totalWeight int) int {
	// From top to bottom. We think about whether to take each one or not,
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

func GetBestValueDP(items []Item, totalWeight int) int {
	// From bottom to top, every best value we are trying to get is based on
	// previous calculated values. Idea is that given maxWeight whether adding
	// this item will result in a greater value than not adding it

	// Create dp mem, which is a 2d slice, size M (number of items) * N (weight)
	var mem [][]int = make([][]int, len(items))
	for i := 0; i < len(items); i++ {
		mem[i] = make([]int, totalWeight+1)
		mem[i][0] = 0
	}

	// For every possible weight, seek the best value that we can get when deciding
	// whether to add the item or not
	for i := 1; i <= totalWeight; i++ {
		// Always add first item to cart if have enough weight
		if items[0].weight <= i {
			mem[0][i] = items[0].value
		}

		for j := 1; j < len(items); j++ {
			// Get the best value of whether to add this item or not
			if items[j].weight <= i {
				mem[j][i] = mem[j-1][i-items[j].weight] + items[j].value
			}

			if mem[j-1][i] > mem[j][i] {
				mem[j][i] = mem[j-1][i]
			}
		}
	}

	return mem[len(items)-1][totalWeight]
}
