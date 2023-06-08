// Input : an array of integers.

// Output : this array, but sorted in such a way that there are two wings:

// the left wing with numbers decreasing,

// the right wing with numbers increasing.

// the two wings have the same length. If the length of the array is odd the wings are around the bottom, if the length is even the
// bottom is considered to be part of the right wing.

// each integer l of the left wing must be greater or equal to its counterpart r in the right wing, the difference l - r being as
// small as possible. In other words the right wing must be nearly as steep as the left wing.
// input --> [79, 35, 54, 19, 35, 25]
// output --> [79, 35, 25, *19*, 35, 54]
// input --> [67, 93, 100, -16, 65, 97, 92]
// output --> [100, 93, 67, *-16*, 65, 92, 97]

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{67, 93, 100, -16, 65, 97, 92}
	sort.Ints(arr)
	reversed_arr := ArrayReverse(arr)
	// fmt.Println(reversed_arr)
	fmt.Println(SortLeftAndRight(reversed_arr))
}

func ArrayReverse(arr []int) []int {
	sortedReversedArray := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		sortedReversedArray = append(sortedReversedArray, arr[i])
	}
	return sortedReversedArray
}

func SortLeftAndRight(arr []int) []int {
	left := []int{}
	right := []int{}

	for ind, i := range arr {
		if ind%2 == 0 {
			left = append(left, i)
		} else {
			right = append(right, i)
		}
	}
	final_result := []int{}
	final_result = append(left, ArrayReverse(right)...)

	return final_result
}
