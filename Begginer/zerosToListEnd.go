// Given an array of numbers. Return the same array with all zero's moved to the end of the array.

package main

import "fmt"

func main() {
	arr := []int{7, 4, 0, 5, 1, 0, 8, 0, 1, 0, 6, 0, 7, 0, 88, 0, 9, 5, 0, 0, 6, 0}
	zeroMoved := zerosToArrayend(arr)
	fmt.Printf("Here is the list with all zero's moved to the end of the list: %d", zeroMoved)
}

func zerosToArrayend(arr []int) []int {
	nonZeros := []int{}
	zeros := []int{}

	for _, i := range arr {
		if i == 0 {
			zeros = append(zeros, i)
		} else {
			nonZeros = append(nonZeros, i)
		}
	}

	nonZeros = append(nonZeros, zeros...)
	return nonZeros
}
