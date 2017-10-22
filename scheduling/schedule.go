package scheduling

import (
	"math/rand"
	"time"
	"github.com/kennhung/ftcScoring/model"
)

const (
	TeamsPerMatch = 4
)

type ScheduleSettings struct {
	MatchesPerTeam int
	Type           string
}

// Creates a random schedule for the given parameters and returns it as a list of matches.
func BuildRandomSchedule(teams []model.Team, schedulesettings ScheduleSettings) ([]model.Match, error) {
	numTeams := len(teams)
	matchesPerTeam := schedulesettings.MatchesPerTeam
	numMatches := matchesPerTeam * numTeams

	// Generate a random permutation of the team ordering to fill into the pre-randomized schedule.
	teamShuffle := rand.Perm(numTeams)
	matches := make([]model.Match, numMatches)
	for _, match := range matches {
		for i := 0; i < numTeams; i = i + 4 {
			var ids [4]int
			var index = i
			for j := 0; j < 4; j++ {
				if index < numMatches {
					ids[j] = teamShuffle[index]
					index++
				} else {
					teamShuffle = rand.Perm(numTeams)
					i = 0
					index = i
					ids[j] = teamShuffle[index]
					index++
				}
			}
			match.Red1 = teams[ids[0]].Id
			match.Red2 = teams[ids[1]].Id
			match.Blue1 = teams[ids[2]].Id
			match.Blue2 = teams[ids[3]].Id
		}
	}

	// Fill in the match times.
	matchIndex := 0
	for i := 0; i < numMatches && matchIndex < numMatches; i++ {
		matches[matchIndex].Time = time.Now()
		matchIndex++
	}

	return matches, nil
}
