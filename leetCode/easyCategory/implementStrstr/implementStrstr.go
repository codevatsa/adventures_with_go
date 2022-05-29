package main

import "fmt"

func strStr(haystack string, needle string) int {
	hayStackLen := len(haystack)
	needleLen := len(needle)
	fmt.Println(hayStackLen, needleLen)
	for i := 0; i < hayStackLen-needleLen+1; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}

// We subtract the difference in length first to ensure we dont cross string limit of haystack while looping.
// Then we extract parts of the string upto length of needle and loop through haystack string to find a match

func main() {
	haystack := "hello does value of needle exiiiist in this haystack ?"
	needle := "ii"
	result := strStr(haystack, needle)
	if result == -1 {
		fmt.Printf("'%s' is not found in '%s'\n", needle, haystack)
	} else {
		fmt.Printf("'%s' is found in '%s' at index: %d\n", needle, haystack, result)
	}
}
