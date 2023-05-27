package main

import "fmt"

func main() {
	arr := []int{7, 4, 3, 9, 1, 0, 8, 12, 45, 77, 89, 66, 54, 78, 23, 54, 73, 19}

	odds, evens := []int{}, []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			odds = append(odds, arr[i])
		} else {
			evens = append(evens, arr[i])
		}
	}

	fmt.Printf("Odd numbers in arr are: %v \n", odds)
	fmt.Printf("Even numbers in arr are: %v \n", evens)

}
