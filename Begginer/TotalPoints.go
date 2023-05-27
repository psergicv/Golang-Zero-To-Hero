// Our football team has finished the championship.

// Our team's match results are recorded in a collection of strings. Each match is represented by a string in the format "x:y", where x is our team's score and y is our opponents score.

// For example: ["3:1", "2:2", "0:1", ...]

// Points are awarded for each match as follows:

// if x > y: 3 points (win)
// if x < y: 0 points (loss)
// if x = y: 1 point (tie)
// We need to write a function that takes this collection and returns the number of points our team (x) got in the championship by the rules given above.

package main

import (
	"fmt"
	"strings"
)

func main() {
	game_list1 := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"} // 30 points
	game_list2 := []string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"} // 10 points
	game_list3 := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "1:3", "1:4", "2:3", "2:4", "3:4"} // 15 points

	fmt.Println(pointCalculator(game_list1))
	fmt.Println(pointCalculator(game_list2))
	fmt.Println(pointCalculator(game_list3))

}

func pointCalculator(games []string) int {
	var points int

	for i := 0; i < len(games); i++ {
		splitted_score := strings.Split(games[i], ":")
		if splitted_score[0] > splitted_score[1] {
			points += 3
		} else if splitted_score[0] == splitted_score[1] {
			points += 1
		}
	}
	return points
}
