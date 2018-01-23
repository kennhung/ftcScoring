package play

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

	fields.Highest = 0
	//TODO add highest
}

// Get length of Rankings
func (rankings Rankings) Len() int {
	return len(rankings)
}

// Helper function to implement the required interface for Sort.
func (rankings Rankings) Less(i, j int) bool {
	a := rankings[i]
	b := rankings[j]

	// Use cross-multiplication to keep it in integer math.
	if a.QP*b.Played == b.QP*a.Played {
		if a.RP*b.Played == b.RP*a.Played {
			if a.Highest*b.Played == b.Highest*a.Played {
				return a.Random > b.Random
			}
			return a.Highest*b.Played > b.Highest*a.Played
		}
		return a.RP*b.Played > b.RP*a.Played
	}
	return a.QP*b.Played > b.QP*a.Played
}

// Helper function to implement the required interface for Sort.
func (rankings Rankings) Swap(i, j int) {
	rankings[i], rankings[j] = rankings[j], rankings[i]
}