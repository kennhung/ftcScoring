package play

type Score struct {
	AutoJewels      int
	AutoCryptobox   int
	CryptoboxKeys   int
	RobotInSafeZone int
	Glyphs          int
	ComRows         int
	ComColumns      int
	ComCiphers      int
	RelicsZ1        int
	RelicsZ2        int
	RelicsZ3        int
	RelicsUpright   int
	RobotBalanced   int
	Penalties       [2]int// 0 for Major, 1 for Minor
	ElimDisq        bool
	Cards           map[string]string
}

type ScoreSummary struct {
	Auto  int
	AutoB int
	Tele  int
	EndG  int
	Pen   int
	Tot   int
}

func (score *Score) Summarize(opponentPenalties [2]int, matchType string) *ScoreSummary {
	summary := new(ScoreSummary)

	// Leave the score at zero if the team was disqualified.
	if score.ElimDisq {
		return summary
	}

	// Calculate autonomous score.
	summary.Auto = 30 * score.AutoJewels
	summary.Auto = summary.Auto + 15*score.AutoCryptobox + 30*score.CryptoboxKeys + 10*score.RobotInSafeZone

	// Calculate teleop score.
	summary.Tele = 2*score.Glyphs
	summary.Tele += 10*score.ComRows+20*score.ComColumns+30*score.ComCiphers

	// Calculate endgame here.
	summary.EndG = 10*score.RelicsZ1+20*score.RelicsZ2+40*score.RelicsZ3+15*score.RelicsUpright+20*score.RobotBalanced

	// Calculate penalty points.
	summary.Pen += opponentPenalties[0]*40+ opponentPenalties[1]*10

	//Total Point
	summary.Tot = summary.Auto + summary.Tele + summary.EndG + summary.Pen

	return summary
}

func NewScore() *Score {
	return &Score{ElimDisq: false}
}

