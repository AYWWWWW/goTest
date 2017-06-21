package main

import "sort"
import "math"
import "fmt"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return -1
	}
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		sum := 0
		for j < k {
			sum = nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if sum > target {
				k--
			} else {
				j++
			}
			result = getCloser(sum, result, target)
		}
	}
	return result
}

func getCloser(num1, num2, target int) int {
	if math.Abs(float64(num1-target)) > math.Abs(float64(num2-target)) {
		return num2
	} else {
		return num1
	}
}

func main() {
	nums := []int{6, -18, -20, -7, -15, 9, 18, 10, 1, -20, -17, -19, -3, -5, -19, 10, 6, -11, 1, -17, -15, 6, 17, -18, -3, 16, 19, -20, -3, -17, -15, -3, 12, 1, -9, 4, 1, 12, -2, 14, 4, -4, 19, -20, 6, 0, -19, 18, 14, 1, -15, -5, 14, 12, -4, 0, -10, 6, 6, -6, 20, -8, -6, 5, 0, 3, 10, 7, -2, 17, 20, 12, 19, -13, -1, 10, -1, 14, 0, 7, -3, 10, 14, 14, 11, 0, -4, -15, -8, 3, 2, -5, 9, 10, 16, -4, -3, -9, -8, -14, 10, 6, 2, -12, -7, -16, -6, 10}
	fmt.Println(threeSumClosest(nums, -52))
}
