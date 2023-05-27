package main

import (
	"fmt"
	"strings"
)

func main() {
	vowels := "aeiouAEIOU"
	text := "This is a test text. You can put here anything you want" // 16 vowels
	var counter int

	for _, letter := range text {
		if strings.Contains(vowels, string(letter)) {
			counter += 1
		}
	}
	fmt.Printf("There are %v vowels in the provided text.", counter)
}
