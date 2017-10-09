package scoring

import "math/rand"

type RankingFields struct {
	QP                int
	RP                int
	MatchPoints       int
	AutoPoints        int
	TelePoints        int
	EndGPoints        int
	Random            float64
	Wins              int
	Losses            int
	Ties              int
	Disqualifications int
	Played            int
	Highest           int
}

type Ranking struct {
	TeamId int
	Rank   int
	RankingFields
}

type Rankings []*Ranking

func (fields *RankingFields) AddScoreSummary(ownScore *ScoreSummary, opponentScore *ScoreSummary, disqualified bool) {
	fields.Played += 1

	if disqualified {
		// Don't award any points.
		fields.Disqualifications += 1
		return
	}

	// Assign Qualification Points and wins/losses/ties.
	if ownScore.Tot > opponentScore.Tot {
		fields.QP += 2
		fields.RP += opponentScore.Tot
		fields.Wins += 1
	} else if ownScore.Tot == opponentScore.Tot {
		fields.QP += 1
		fields.RP += opponentScore.Tot
		fields.Ties += 1
	} else {
		fields.RP += ownScore.Tot
		fields.Losses += 1
	}

	// Assign tiebreaker points.
	fields.MatchPoints += ownScore.Tot
	fields.AutoPoints += ownScore.Auto
	fields.TelePoints += ownScore.Tele
	fields.EndGPoints += ownScore.EndG

	// Store a random value to be used as the last tiebreaker if necessary.
	fields.Random = rand.Float64()
}

// Get length of Rankings
func (rankings Rankings) Len() int {
	return len(rankings)
}

// Exchange two Ranking
func (rankings Rankings) Exchange(i, j int) {
	rankings[i], rankings[j] = rankings[j], rankings[i]
}
