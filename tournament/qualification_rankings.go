package tournament

import (

	"sort"
	"strconv"
	"github.com/kennhung/ftcScoring/model"
	"github.com/kennhung/ftcScoring/play"
)

// Determines the rankings from the stored match results, and saves them to the database.
func CalculateRankings(database *model.Database) error {
	matches, err := database.GetMatchesByType("qualification")
	if err != nil {
		return err
	}
	rankings := make(map[int]*play.Ranking)
	for _, match := range matches {
		if match.Status != "complete" {
			continue
		}
		matchResult, err := database.GetMatchResultForMatch(match.Id)
		if err != nil {
			return err
		}
		if !match.Red1notshow {
			addMatchResultToRankings(rankings, match.Red1, matchResult, true)
		}
		if !match.Red2notshow {
			addMatchResultToRankings(rankings, match.Red2, matchResult, true)
		}
		if !match.Blue1notshow {
			addMatchResultToRankings(rankings, match.Blue1, matchResult, false)
		}
		if !match.Blue2notshow {
			addMatchResultToRankings(rankings, match.Blue2, matchResult, false)
		}
	}

	sortedRankings := sortRankings(rankings)
	for rank, ranking := range sortedRankings {
		ranking.Rank = rank + 1
	}
	err = database.ReplaceAllRankings(sortedRankings)
	if err != nil {
		return nil
	}

	return nil
}

// Checks all the match results for yellow and red cards, and updates the team model accordingly.
func CalculateTeamCards(database *model.Database, matchType string) error {
	teams, err := database.GetAllTeams()
	if err != nil {
		return err
	}
	teamsMap := make(map[string]model.Team)
	for _, team := range teams {
		team.YellowCard = false
		teamsMap[strconv.Itoa(team.Id)] = team
	}

	matches, err := database.GetMatchesByType(matchType)
	if err != nil {
		return err
	}
	for _, match := range matches {
		if match.Status != "complete" {
			continue
		}
		matchResult, err := database.GetMatchResultForMatch(match.Id)
		if err != nil {
			return err
		}

		// Mark the team as having a yellow card if they got either a yellow or red in a previous match.
		for teamId, card := range matchResult.RedCards {
			if team, ok := teamsMap[teamId]; ok && card != "" {
				team.YellowCard = true
				teamsMap[teamId] = team
			}
		}
		for teamId, card := range matchResult.BlueCards {
			if team, ok := teamsMap[teamId]; ok && card != "" {
				team.YellowCard = true
				teamsMap[teamId] = team
			}
		}
	}

	// Save the teams to the database.
	for _, team := range teamsMap {
		err = database.SaveTeam(&team)
		if err != nil {
			return err
		}
	}

	return nil
}

// Incrementally accounts for the given match result in the set of rankings that are being built.
func addMatchResultToRankings(rankings map[int]*play.Ranking, teamId int, matchResult *model.MatchResult, isRed bool) {
	ranking := rankings[teamId]
	if ranking == nil {
		ranking = &play.Ranking{TeamId: teamId}
		rankings[teamId] = ranking
	}

	// Determine whether the team was disqualified.
	var cards map[string]string
	if isRed {
		cards = matchResult.RedCards
	} else {
		cards = matchResult.BlueCards
	}
	disqualified := false
	if card, ok := cards[strconv.Itoa(teamId)]; ok && card == "red" {
		disqualified = true
	}



	if isRed {
		ranking.AddScoreSummary(matchResult.RedScoreSummary(), matchResult.BlueScoreSummary(), disqualified)
	} else {
		ranking.AddScoreSummary(matchResult.BlueScoreSummary(), matchResult.RedScoreSummary(), disqualified)
	}
}

func sortRankings(rankings map[int]*play.Ranking) play.Rankings {
	var sortedRankings play.Rankings
	for _, ranking := range rankings {
		sortedRankings = append(sortedRankings, ranking)
	}
	sort.Sort(sortedRankings)
	return sortedRankings
}
