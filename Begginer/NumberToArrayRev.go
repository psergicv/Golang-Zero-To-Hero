// Given a random non-negative number, you have to return the digits of this number within an slice in reverse order.

// Example(Input => Output):
// 35231 => [1,3,2,5,3]
// 0 => [0]

package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := 45786
	fmt.Println(ReversedNumberToSlice(num1))
}

func ReversedNumberToSlice(num int) []int {
	str_num := strconv.Itoa(num)
	var arr []int

	for i := len(str_num) - 1; i >= 0; i-- {
		sn := string(str_num[i])
		int_sn, err := strconv.Atoi(sn)
		if err != nil {
			fmt.Println(err)
		} else {
			arr = append(arr, int_sn)
		}
	}
	return arr
}
