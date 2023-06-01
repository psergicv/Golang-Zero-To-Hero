// Trolls are attacking your comment section!

// A common way to deal with this situation is to remove all of the vowels from the trolls' comments, neutralizing the threat.

// Your task is to write a function that takes a string and return a new string with all vowels removed.

// For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".

// Note: for this kata y isn't considered a vowel.

package main

import (
	"fmt"
	"strings"
)

func main() {
	word := Disevowel("This website is for losers LOL!")
	fmt.Printf("The divoweled sentence is: %v", word)
}

func Disevowel(comment string) string {
	vowels := "aeiouAEIOU"
	result := []string{}

	for i := 0; i < len(comment); i++ {
		if !strings.Contains(vowels, string(comment[i])) {
			result = append(result, string(comment[i]))
		}
	}
	return strings.Join(result, "")
}
