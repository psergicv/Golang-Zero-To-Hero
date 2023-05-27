package main

import "fmt"

func main() {
	arr1 := []int{7, 1, 5, 4, 6, 3, 2, 9}
	arr2 := []int{5, 3, 4, 0, 9, 4, 5, 1}
	commonElements := []int{}

	for _, element := range arr1 {
		if contains(arr2, element) {
			commonElements = append(commonElements, element)
		}
	}

	fmt.Printf("Common elements between arr1 and arr2 are: %v", commonElements)

}

func contains(s []int, searchElement int) bool {
	for _, i := range s {
		if i == searchElement {
			return true
		}
	}
	return false
}
