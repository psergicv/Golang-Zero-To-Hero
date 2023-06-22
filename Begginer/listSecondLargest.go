// Given an unsorted array of numbers. Return the second larges number of the array.

package main

import (
	"fmt"
)

func main() {
	arr := []int{7, 8, 45, 199, 2020}
	sl := secondLargest(arr)
	fmt.Printf("The second larges number is: %d\n", sl)
}

func secondLargest(arr []int) int {
	first, second := 0, 0
	for _, i := range arr {
		if i > second && i > first {
			second = first
			first = i
		} else if second < i && i <= first {
			second = i
		}
	}
	return second
}
