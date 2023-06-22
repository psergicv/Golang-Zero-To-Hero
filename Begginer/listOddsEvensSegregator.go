// Given an array arr[] of integers, segregate even and odd numbers in the array. Such that all the even numbers should be
// present first, and then the odd numbers.
// Examples:

// Input: arr[] = 1 9 5 3 2 6 7 11
// Output: 2 6 5 3 1 9 7 11

// Input: arr[] = 1 3 2 4 7 6 9 10
// Output:  2 4 6 10 7 1 9 3

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 9, 5, 3, 2, 6, 7, 11}
	segregatedNums := oddEvenSegregator(nums)
	fmt.Printf("Segregated array of nums is: %d\n", segregatedNums)
}

func oddEvenSegregator(nums []int) []int {
	odds := []int{}
	even := []int{}

	for _, i := range nums {
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odds = append(odds, i)
		}
	}
	combinedList := append(even, odds...)

	return combinedList
}
