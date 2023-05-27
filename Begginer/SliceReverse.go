package main

import "fmt"

func main() {
	s1 := []int{7, 9, 3, 4, 5, 0}
	s2 := []string{"Python", "Go", "Ruby", "JS"}

	fmt.Println(reverseStringSlice(s2))
	fmt.Println(reverseIntSlice(s1))

}

func reverseStringSlice(s []string) []string {
	new_slice := []string{}

	for i := len(s) - 1; i >= 0; i-- {
		new_slice = append(new_slice, s[i])
	}
	return new_slice
}

func reverseIntSlice(s []int) []int {
	new_slice := []int{}

	for i := len(s) - 1; i >= 0; i-- {
		new_slice = append(new_slice, s[i])
	}
	return new_slice
}
