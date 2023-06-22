// Rearrange array such that even positioned are greater than odd

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 2, 1}
	rearangedArray := odds_evens_inds(nums)
	fmt.Printf("The rearranged array is: %d\n", rearangedArray)
}

func odds_evens_inds(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		} else {
			if nums[i] < nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums
}
