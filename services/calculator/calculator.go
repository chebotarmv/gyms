package calculator

import "sort"

// FindPerfectPackets finds the perfect combination of numbers from nums that add up to target.
// Exact Combination Search: O(n)
// Closest Combination Search (with recursion): O(2^n) in the worst case, but can be significantly lower in most practical cases due to constraints.
// Overall Algorithmic Complexity: In the worst case, when searching for the closest combination using recursion, the overall complexity would be O(n + 2^n).
// Overall, the total memory usage for this function will be approximately O(n) for input data and O(n) for the recursion call stack,
// which in the worst case gives O(n) + O(n) = O(n) additional memory.
func FindPerfectPackets(availablePacketSizes []int, itemsToStore int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(availablePacketSizes)))

	// try to find an exact match
	exactTarget := itemsToStore
	var exactResult []int
	for _, num := range availablePacketSizes {
		for exactTarget >= num {
			exactResult = append(exactResult, num)
			exactTarget -= num
		}
	}

	// if we found an exact match, return it
	if exactTarget == 0 {
		return exactResult
	}

	// otherwise, try to find the closest match
	var closestCombination []int
	var closestDiff int

	var backtrack func(index int, currentSum int, currentCombination []int)
	backtrack = func(index int, currentSum int, currentCombination []int) {
		diff := currentSum - itemsToStore
		if diff > 0 && (closestCombination == nil || diff < closestDiff) {
			closestCombination = append([]int{}, currentCombination...)
			closestDiff = diff
		}

		if index >= len(availablePacketSizes) || currentSum >= itemsToStore {
			return
		}

		backtrack(index, currentSum+availablePacketSizes[index], append(currentCombination, availablePacketSizes[index]))
		backtrack(index+1, currentSum, currentCombination)
	}

	backtrack(0, 0, []int{})
	return closestCombination
}
