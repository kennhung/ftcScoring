package model

import (
	"encoding/json"
	"github.com/kennhung/ftcScoring/scoring"
)

type MatchResult struct {
	Id         int
	MatchId    int
	PlayNumber int
	MatchType  string
	RedScore   *scoring.Score
	BlueScore  *scoring.Score
	RedCards   map[string]string
	BlueCards  map[string]string
}

type MatchResultDb struct {
	Id            int
	MatchId       int
	PlayNumber    int
	MatchType     string
	RedScoreJson  string
	BlueScoreJson string
	RedCardsJson  string
	BlueCardsJson string
}

// Returns a new match result object with empty slices instead of nil.
func NewMatchResult() *MatchResult {
	matchResult := new(MatchResult)
	matchResult.RedScore = new(scoring.Score)
	matchResult.BlueScore = new(scoring.Score)
	matchResult.RedCards = make(map[string]string)
	matchResult.BlueCards = make(map[string]string)
	return matchResult
}

func (database *Database) CreateMatchResult(matchResult *MatchResult) error {
	matchResultDb, err := matchResult.Serialize()
	if err != nil {
		return err
	}
	err = database.matchResultMap.Insert(matchResultDb)
	if err != nil {
		return err
	}
	matchResult.Id = matchResultDb.Id
	return nil
}

func (database *Database) GetMatchResultForMatch(matchId int) (*MatchResult, error) {
	var matchResults []MatchResultDb
	query := "SELECT * FROM match_results WHERE matchid = ? ORDER BY playnumber DESC LIMIT 1"
	err := database.matchResultMap.Select(&matchResults, query, matchId)
	if err != nil {
		return nil, err
	}
	if len(matchResults) == 0 {
		return nil, nil
	}
	matchResult, err := matchResults[0].Deserialize()
	if err != nil {
		return nil, err
	}
	return matchResult, err
}

func (database *Database) SaveMatchResult(matchResult *MatchResult) error {
	matchResultDb, err := matchResult.Serialize()
	if err != nil {
		return err
	}
	_, err = database.matchResultMap.Update(matchResultDb)
	return err
}

func (database *Database) DeleteMatchResult(matchResult *MatchResult) error {
	matchResultDb, err := matchResult.Serialize()
	if err != nil {
		return err
	}
	_, err = database.matchResultMap.Delete(matchResultDb)
	return err
}

func (database *Database) TruncateMatchResults() error {
	return database.matchResultMap.TruncateTables()
}

// Calculates and returns the summary fields used for ranking and display for the red alliance.
func (matchResult *MatchResult) RedScoreSummary() *scoring.ScoreSummary {
	return matchResult.RedScore.Summarize(matchResult.BlueScore.Penalties, matchResult.MatchType)
}

// Calculates and returns the summary fields used for ranking and display for the blue alliance.
func (matchResult *MatchResult) BlueScoreSummary() *scoring.ScoreSummary {
	return matchResult.BlueScore.Summarize(matchResult.RedScore.Penalties, matchResult.MatchType)
}

// Checks the score for disqualifications or a tie and adjusts it appropriately.
func (matchResult *MatchResult) CorrectEliminationScore() {
	matchResult.RedScore.ElimDisq = false
	for _, card := range matchResult.RedCards {
		if card == "red" {
			matchResult.RedScore.ElimDisq = true
		}
	}
	for _, card := range matchResult.BlueCards {
		if card == "red" {
			matchResult.BlueScore.ElimDisq = true
		}
	}

	// No elimination tiebreakers.
}

// Converts the nested struct MatchResult to the DB version that has JSON fields.
func (matchResult *MatchResult) Serialize() (*MatchResultDb, error) {
	matchResultDb := MatchResultDb{Id: matchResult.Id, MatchId: matchResult.MatchId,
		PlayNumber: matchResult.PlayNumber, MatchType: matchResult.MatchType}
	if err := toJson(&matchResultDb.RedScoreJson, matchResult.RedScore); err != nil {
		return nil, err
	}
	if err := toJson(&matchResultDb.BlueScoreJson, matchResult.BlueScore); err != nil {
		return nil, err
	}
	if err := toJson(&matchResultDb.RedCardsJson, matchResult.RedCards); err != nil {
		return nil, err
	}
	if err := toJson(&matchResultDb.BlueCardsJson, matchResult.BlueCards); err != nil {
		return nil, err
	}
	return &matchResultDb, nil
}

// Converts the DB MatchResult with JSON fields to the nested struct version.
func (matchResultDb *MatchResultDb) Deserialize() (*MatchResult, error) {
	matchResult := MatchResult{Id: matchResultDb.Id, MatchId: matchResultDb.MatchId,
		PlayNumber: matchResultDb.PlayNumber, MatchType: matchResultDb.MatchType}
	if err := json.Unmarshal([]byte(matchResultDb.RedScoreJson), &matchResult.RedScore); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(matchResultDb.BlueScoreJson), &matchResult.BlueScore); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(matchResultDb.RedCardsJson), &matchResult.RedCards); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(matchResultDb.BlueCardsJson), &matchResult.BlueCards); err != nil {
		return nil, err
	}
	return &matchResult, nil
}
