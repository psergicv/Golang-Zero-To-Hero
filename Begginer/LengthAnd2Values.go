// Given an integer n and two other values, build an array of size n filled with these two values alternating.

// 5, true, false     -->  [true, false, true, false, true]
// 10, "blue", "red"  -->  ["blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"]
// 0, "one", "two"    -->  []

package main

import "fmt"

func main() {
	res := LenVal(10, "blue", "red")
	fmt.Println(res)
}

func LenVal(n int, s1 string, s2 string) []string {
	if n == 0 {
		return []string{}
	} else {
		var array_result = []string{}
		var i int = 0
		for i != n {
			values := [2]string{s1, s2}
			if i%2 != 0 {
				array_result = append(array_result, string(values[1]))
			} else {
				array_result = append(array_result, string(values[0]))
			}
			i += 1
		}
		return array_result
	}
}
