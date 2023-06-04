// The Commentator's Insight

// Objective: You are a commentator and have access to the season's goal data for a set of teams. The goal data is stored as a map
// where the key is a team name and the value is another map (with keys 'home' and 'away' for home and away goals). Your task is
// to write a function that accepts the map and a team name, and returns the total number of goals scored by the team in the season
// (both home and away).

package main

import (
	"fmt"
)

func main() {
	leagueResults := map[string]map[string]int{
		"Real Madrid": map[string]int{"home": 30, "away": 20},
		"Barcelona":   map[string]int{"home": 28, "away": 22},
		"Juventus":    map[string]int{"home": 25, "away": 18},
	}
	result := CommInsight(leagueResults)

	for team, goals := range result {
		fmt.Printf("%v has scored %v goals this season\n", team, goals)
	}
}

func CommInsight(ResultDB map[string]map[string]int) map[string]int {
	var teamGoals = make(map[string]int)
	for team_name, team := range ResultDB {
		goalCounter := 0
		for _, goalScored := range team {
			goalCounter += int(goalScored)
			teamGoals[team_name] = goalCounter
		}
	}
	return teamGoals
}
