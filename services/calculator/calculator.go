package calculator

import "sort"

// FindPerfectFit finds the perfect combination of numbers from nums that add up to target.
func FindPerfectFit(nums []int, target int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	// try to find an exact match
	exactTarget := target
	var exactResult []int
	for _, num := range nums {
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
		diff := currentSum - target
		if diff > 0 && (closestCombination == nil || diff < closestDiff) {
			closestCombination = append([]int{}, currentCombination...)
			closestDiff = diff
		}

		if index >= len(nums) || currentSum >= target {
			return
		}

		backtrack(index, currentSum+nums[index], append(currentCombination, nums[index]))
		backtrack(index+1, currentSum, currentCombination)
	}

	backtrack(0, 0, []int{})
	return closestCombination
}
