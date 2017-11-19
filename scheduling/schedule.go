package scheduling

import (
	"math/rand"
	"time"
	"github.com/kennhung/ftcScoring/model"
	"fmt"
	"math"
)

const (
	TeamsPerMatch = 4
)


// Creates a random schedule for the given parameters and returns it as a list of matches.
func BuildRandomSchedule(teams []model.Team, MatchesPerTeam int, Type string) ([]model.Match, error) {
	numTeams := len(teams)
	matchesPerTeam := MatchesPerTeam
	numMatches := int(math.Ceil(float64(matchesPerTeam * numTeams)/4.0))
	fmt.Print(math.Ceil(float64(matchesPerTeam * numTeams)/4.0))

	// Generate a random permutation of the team ordering to fill into the pre-randomized schedule.
	teamShuffle := rand.Perm(numTeams)
	matches := make([]model.Match, numMatches)

	i := 0
	var ids [4]int
	index := i

	for num, match := range matches {
		for j := 0; j < 4; j++ {
			if index < numTeams {
				ids[j] = teamShuffle[index];
				index++;
			} else {
				teamShuffle = rand.Perm(numTeams)
				i = 0
				index = i
				ids[j] = teamShuffle[index]
				index++
			}
		}
		match.Red1 = teams[ids[0]].Id;
		match.Red2 = teams[ids[1]].Id;
		match.Blue1 = teams[ids[2]].Id;
		match.Blue2 = teams[ids[3]].Id;
		match.Type = Type
		match.DisplayName = fmt.Sprint(num+1)
		i = i + 4
		matches[num] = match
	}

	// Fill in the match times.
	matchIndex := 0
	for i := 0; i < numMatches && matchIndex < numMatches; i++ {
		matches[matchIndex].Time = time.Now()
		matchIndex++
	}

	return matches, nil
}
