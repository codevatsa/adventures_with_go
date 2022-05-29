package main

import "fmt"

// Brute force - run through array once and store it in hash. Run through len of array once and check if that key exists in hash
func MissingNumber(nums []int) int {
	mapOfNum := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		mapOfNum[nums[i]] = i
	}
	for j := 1; j < n+1; j++ {
		if _, ok := mapOfNum[j]; !ok {
			return j
		}
	}
	return 0 // No numbers are missing in this case
}

// This method uses sum of n formula. In an array of n numbers - sum of elements of that array will be n*(n+1)/2 .
// If all numbers from 0 to n exists in the array and when we subtract all entries from the expected sum -> value should be zero
func missingNumberSumMethod(nums []int) int {
	n := len(nums)
	sumOfN := (n * (n + 1)) / 2
	for _, val := range nums {
		sumOfN -= val
	}
	return sumOfN
}

// Slightly complex version using XOR method of bit manipulation. XOR of a number with itself is 0. This can be used to eliminate existing elements
// and the last val that remains after XOR ing all the elements of the array will be the  missing number
func missingNumberXorMethod(nums []int) int {
	res := 0
	for i, val := range nums {
		res ^= val
		res ^= i + 1
	}
	return res
}

func main() {
	var inputArr []int = []int{9, 6, 4, 2, 3, 5, 7, 0, 1} // 8 is missing in this array

	res1 := MissingNumber(inputArr)
	fmt.Println("Result hashmap method: ", res1)

	res2 := missingNumberSumMethod(inputArr)
	fmt.Println("Result optimised - sum of array: ", res2)

	res3 := missingNumberXorMethod(inputArr)
	fmt.Println("Result optimised - XOR method: ", res3)
}
