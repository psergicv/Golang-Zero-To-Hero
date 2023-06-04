// Task 1: The Scout's List

// Objective: You are a scout who has travelled to watch a match with the aim of recruiting a striker. You've written down the names
// of all the strikers playing the match in a list. Your task is to create a function that counts how many times each player has
// touched the ball. Use a dictionary where keys are player names and values are the number of touches.

package main

import (
	"fmt"
	"strings"
)

func main() {
	player_list := []string{"Mbappe", "Lewandowski", "Cavani", "Lewandowski", "Lewandowski", "Mbappe"}
	final_list := ScoutForwardList(player_list)
	fmt.Printf("Here is how many times each player tocuhed the ball: %v", final_list)
}

func ScoutForwardList(playerList []string) map[string]int {
	var final_list = make(map[string]int)

	for _, player := range playerList {
		final_list[player] = strings.Count(strings.Join(playerList, " "), player)
	}
	return final_list
}
