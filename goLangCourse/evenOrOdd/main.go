package main

import "fmt"

func main() {
	// 1st req: Create slice of int from 0 through 10
	numsToCheck := generateNumberSlice(10)
	for _, val := range numsToCheck {
		if val%2 == 0 {
			fmt.Printf("%v is even \n", val)
		} else {
			fmt.Printf("%v is odd \n", val)
		}
	}
}

func generateNumberSlice(size int) []int {
	sliceToSend := []int{}
	for i := 1; i <= size; i++ {
		sliceToSend = append(sliceToSend, i)
	}
	return sliceToSend
}
