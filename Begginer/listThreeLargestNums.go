// Given an unsorted list of numbers. Return the last 3 largest numbers from the list.
package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 7, 8, 9, 12, 33}
	f, s, t := threeLargestNums(arr)
	fmt.Printf("First largest number is: %d\n", f)
	fmt.Printf("Second largest number is: %d\n", s)
	fmt.Printf("Third largest number is: %d\n", t)
}

func threeLargestNums(arr []int) (int, int, int) {
	first, second, third := 0, 0, 0

	for _, i := range arr {
		if i >= first {
			first = i
		}
	}
	for _, i := range arr {
		if second <= i && i < first {
			second = i
		}
	}
	for _, i := range arr {
		if third <= i && (i < second && i < first) {
			third = i
		}
	}
	return first, second, third
}
