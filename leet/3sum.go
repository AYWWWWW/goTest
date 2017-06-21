package main

import "sort"

func threeSum(nums []int) [][]int {
	var results [][]int
	if len(nums) < 3 {
		return results
	}
	sort.Ints(nums)
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				results = append(results, []int{nums[i], nums[j], nums[k]})
				for nums[j] == nums[j+1] && j < k {
					j++
				}
				for nums[k] == nums[k-1] && j < k {
					k--
				}
			} else if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			}
		}
	}
	return results
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	threeSum(nums)
}
