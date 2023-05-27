package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Golang programming"
	fmt.Println(reverseString(word))
}

func reverseString(s string) string {
	var temp []string
	for i := len(s) - 1; i >= 0; i-- {
		temp = append(temp, string(s[i]))
	}
	reversed_str := strings.Join(temp, "")

	return reversed_str
}
