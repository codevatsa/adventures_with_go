package main

import (
	"fmt"
)

// BruteForce - In the worst case - time complexity of O(n^2), space complexity of O(1)
func TwoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// Using hashmap - store keys which we have visited already and then check if target - value exists in the hashmap
// time complexity is O(n) and space complexity is O(n) since we need to check all elements atleast once in the worst case
func twoSumOptimsed(nums []int, target int) []int {
	mapOfIndex := make(map[int]int)
	for i, currentVal := range nums {
		if reqIndex, valExists := mapOfIndex[target-currentVal]; valExists {
			return []int{reqIndex, i}
		}
		mapOfIndex[currentVal] = i
	}
	return []int{}
}

func main() {
	var inputArr []int = []int{1, 3, 11, 12, 15, 16, 17, 22, 26, 36, 54, 57, 59, 65, 69, 72, 74, 79, 80, 90, 91, 102, 105,
		107, 112, 119, 124, 129, 131, 136, 137, 146, 147, 161, 168, 169, 178, 183, 184, 186, 191, 198, 206, 207, 208, 212,
		216, 218, 219, 224, 226, 230, 234, 237, 239, 241, 243, 246, 253, 256, 257, 258, 259, 261, 262, 265, 267, 276, 277,
		280, 282, 284, 288, 291, 292, 293, 294, 296, 308, 309, 326, 328, 344, 353, 355, 361, 385, 393, 397, 401, 432, 437, 439, 443, 462, 465, 485, 488, 494, 500}
	targetSum := 988 // 488 + 500, i.e 97 and 99th indices

	res1 := TwoSum(inputArr, targetSum)
	fmt.Println("Result brute force: ", res1)

	res2 := twoSumOptimsed(inputArr, targetSum)
	fmt.Println("Result optimised: ", res2)
}
