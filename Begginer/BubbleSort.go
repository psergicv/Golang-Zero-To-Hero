package main

import "fmt"

func main() {
	test := []int{7, 5, 9, 2, 3, 4, 1, 8}
	sorted := BubbleSort(test)
	fmt.Println(sorted)
}

func BubbleSort(arr []int) []int {
	l := 0
	for l < len(arr) {
		for i := 0; i <= len(arr)-2; i++ {
			if arr[i] >= arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		l++
	}
	return arr
}
